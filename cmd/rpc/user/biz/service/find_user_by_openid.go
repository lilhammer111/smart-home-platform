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

type FindUserByOpenidService struct {
	ctx context.Context
}

// NewFindUserByOpenidService new FindUserByOpenidService
func NewFindUserByOpenidService(ctx context.Context) *FindUserByOpenidService {
	return &FindUserByOpenidService{ctx: ctx}
}

// Run create note info
func (s *FindUserByOpenidService) Run(req *micro_user.RpcFindUserByOpenidReq) (resp *user.UserInfo, err error) {
	userInfo := model.User{}
	if err = db.GetMysql().Where("openid = ?", req.Openid).First(&userInfo).Error; err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, fmt.Sprintf("acount of openid %s do not found", req.Openid))
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &user.UserInfo{}
	if err = copier.Copy(resp, &userInfo); err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return resp, nil
}
