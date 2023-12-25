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
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FindUserService struct {
	ctx context.Context
} // NewFindUserService new FindUserService
func NewFindUserService(ctx context.Context) *FindUserService {
	return &FindUserService{
		ctx: ctx,
	}
}

// Run create note info
func (s *FindUserService) Run(req *micro_user.RpcFindUserReq) (resp *user.UserInfo, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().First(&userInfo, req.UserId).Error; err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, fmt.Sprintf("account %d do not exist", req.UserId))
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	userInfo.Password = nil

	resp = &user.UserInfo{}
	if err = copier.Copy(resp, &userInfo); err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return
}
