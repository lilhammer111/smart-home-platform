package service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UserInfo) (resp *user.UserInfo, err error) {
	userInfo := model.User{}
	userInfo.ID = *req.Id
	res := db.GetMysql().Model(&userInfo).Updates(req)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, bizerr.NewInternalError(err)
	}
	if res.RowsAffected == 0 {
		klog.Info("update nothing")
		return nil, kerrors.NewBizStatusError(code.NotFound, "nothing was updated")
	}

	resp = &user.UserInfo{}
	if err = copier.Copy(resp, &userInfo); err != nil {
		klog.Error(err)
		return nil, err
	}

	return resp, nil
}
