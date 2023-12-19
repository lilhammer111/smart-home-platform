package service

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FindUserByUsernameService struct {
	ctx context.Context
} // NewFindUserByUsernameService new FindUserByUsernameService
func NewFindUserByUsernameService(ctx context.Context) *FindUserByUsernameService {
	return &FindUserByUsernameService{ctx: ctx}
}

// Run create note info
func (s *FindUserByUsernameService) Run(req *micro_user.RpcFindUserByUsernameReq) (resp *user.UserInfo, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().Where("username = ?", req.Username).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewInternalError(err)
	}

	if err = copier.Copy(resp, &userInfo); err != nil {
		return nil, bizerr.NewInternalError(err)
	}

	return
}
