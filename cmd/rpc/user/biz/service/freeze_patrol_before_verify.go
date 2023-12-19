package service

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FreezePatrolBeforeVerifyService struct {
	ctx context.Context
}

// NewFreezePatrolBeforeVerifyService new FreezePatrolBeforeVerifyService
func NewFreezePatrolBeforeVerifyService(ctx context.Context) *FreezePatrolBeforeVerifyService {
	return &FreezePatrolBeforeVerifyService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolBeforeVerifyService) Run(req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	userInfo := model.User{}
	// todo To consider the case where the client is passing in multiple conditions at the same time
	err = db.GetMysql().
		Select("is_frozen", "thawed_at").
		Where("mobile = ?", req.Mobile).Or("username = ?", req.Username).
		Or("email = ?", req.Email).First(&userInfo).Error
	if err != nil {
		klog.Errorf("failed to query user model: %s", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}
	if err = copier.Copy(resp, &userInfo); err != nil {
		klog.Errorf("failed to copy user model stuct to RpcFreezeResp struct: %s", err)
		return nil, bizerr.NewInternalError(err)
	}
	return
}
