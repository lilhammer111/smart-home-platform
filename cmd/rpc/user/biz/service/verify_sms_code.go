package service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &micro_user.RpcVerifyResp{}
	if err = bcrypt.CompareHashAndPassword([]byte(hashedSmsCode), []byte(req.SmsCode)); err != nil {
		klog.Error(err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, kerrors.NewBizStatusError(code.AuthenticationFailed, bcrypt.ErrMismatchedHashAndPassword.Error())
		}
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	resp.VerifyPassed = true
	return resp, nil
}
