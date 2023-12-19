package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
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
	if err = db.GetMysql().Model(&userInfo).Updates(req).Error; err != nil {
		return nil, bizerr.NewInternalError(err)
	}
	return
}
