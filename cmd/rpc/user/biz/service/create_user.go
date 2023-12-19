package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/jinzhu/copier"
)

type CreateUserService struct {
	ctx context.Context
} // NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateUserService) Run(req *user.UserInfo) (resp *user.UserInfo, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().Create(&userInfo).Error; err != nil {
		return nil, bizerr.NewInternalErr(err)
	}
	err = copier.Copy(resp, &userInfo)
	if err != nil {
		return nil, bizerr.NewInternalErr(err)
	}
	return
}
