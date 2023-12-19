package service

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type VerifyEmailPwdService struct {
	ctx context.Context
} // NewVerifyEmailPwdService new VerifyEmailPwdService
func NewVerifyEmailPwdService(ctx context.Context) *VerifyEmailPwdService {
	return &VerifyEmailPwdService{ctx: ctx}
}

// Run create note info
func (s *VerifyEmailPwdService) Run(req *micro_user.RpcVerifyEmailPwdReq) (resp *micro_user.RpcVerifyResp, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().Where("email = ?", req.Email).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}

	resp = &micro_user.RpcVerifyResp{}
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.EntryPwd))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, bizerr.NewAuthenticationError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}

	resp.VerifyPassed = true
	return resp, nil
}
