package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	common "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(userId int32) (resp *common.Empty, err error) {
	if err = db.GetMysql().Delete(model.User{}, userId).Error; err != nil {
		return nil, bizerr.NewInternalErr(err)
	}
	return &common.Empty{}, nil
}
