package service

import (
	"context"
	"errors"
	"fmt"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, fmt.Sprintf("account of email %s is not found", req.Email))
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &micro_user.RpcVerifyResp{}
	err = bcrypt.CompareHashAndPassword([]byte(*userInfo.Password), []byte(req.EntryPwd))
	if err != nil {
		klog.Error(err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return resp, kerrors.NewBizStatusError(code.AuthenticationFailed, bcrypt.ErrMismatchedHashAndPassword.Error())
		}
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	resp.VerifyPassed = true
	return resp, nil
}
