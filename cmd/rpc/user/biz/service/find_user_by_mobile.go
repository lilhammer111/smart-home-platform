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

type FindUserByMobileService struct {
	ctx context.Context
} // NewFindUserByMobileService new FindUserByMobileService
func NewFindUserByMobileService(ctx context.Context) *FindUserByMobileService {
	return &FindUserByMobileService{ctx: ctx}
}

// Run create note info
func (s *FindUserByMobileService) Run(req *micro_user.RpcFindUserByMobileReq) (resp *user.UserInfo, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().Where("mobile = ?", req.Mobile).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Error(err)
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}

	resp = &user.UserInfo{}
	if err = copier.Copy(resp, &userInfo); err != nil {
		klog.Error(err)
		return nil, bizerr.NewInternalError(err)
	}
	return
}
