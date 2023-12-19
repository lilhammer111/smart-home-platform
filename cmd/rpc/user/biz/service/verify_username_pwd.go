package service

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type VerifyUsernamePwdService struct {
	ctx context.Context
} // NewVerifyUsernamePwdService new VerifyUsernamePwdService
func NewVerifyUsernamePwdService(ctx context.Context) *VerifyUsernamePwdService {
	return &VerifyUsernamePwdService{ctx: ctx}
}

// Run create note info
func (s *VerifyUsernamePwdService) Run(req *micro_user.RpcVerifyUsernamePwdReq) (resp *micro_user.RpcVerifyResp, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().Where("username = ?", req.Username).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}

	resp = &micro_user.RpcVerifyResp{}
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.EntryPwd))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return resp, bizerr.NewAuthenticationError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}

	resp.VerifyPassed = true
	return resp, nil
}
