package service

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
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
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewExternalError(err)
	}
	resp = &user.UserInfo{}
	if err = copier.Copy(resp, &userInfo); err != nil {
		klog.Error(err)
		return nil, bizerr.NewInternalError(err)
	}
	return
}
