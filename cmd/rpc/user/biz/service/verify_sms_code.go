package service

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"golang.org/x/crypto/bcrypt"
)

type VerifySmsCodeService struct {
	ctx context.Context
} // NewVerifySmsCodeService new VerifySmsCodeService
func NewVerifySmsCodeService(ctx context.Context) *VerifySmsCodeService {
	return &VerifySmsCodeService{ctx: ctx}
}

// Run create note info
func (s *VerifySmsCodeService) Run(req *micro_user.RpcVerifyCodeReq) (resp *micro_user.RpcVerifyResp, err error) {
	hashedSmsCode, err := db.GetRedis().Get(context.Background(), req.Mobile).Result()
	if err != nil {
		return nil, bizerr.NewExternalError(err)
	}

	resp = &micro_user.RpcVerifyResp{}
	if err = bcrypt.CompareHashAndPassword([]byte(hashedSmsCode), []byte(req.SmsCode)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, bizerr.NewAuthenticationError(bcrypt.ErrMismatchedHashAndPassword)
		}
		return nil, bizerr.NewInternalError(err)
	}

	resp.VerifyPassed = true
	return resp, nil
}
