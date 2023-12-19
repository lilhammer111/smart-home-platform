package service

import (
	"context"
	"errors"
	"fmt"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
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
func (s *FreezePatrolBeforeVerifyService) Run(req *micro_user.RpcFreezeReq) (resp *micro_user.RpcUserId, err error) {
	userInfo := model.User{}
	resp = &micro_user.RpcUserId{}
	// todo To consider the case where the client is passing in multiple conditions at the same time
	err = db.GetMysql().
		Select("id").
		Where("mobile = ?", req.Mobile).Or("username = ?", req.Username).
		Or("email = ?", req.Email).First(&userInfo).Error
	if err != nil {
		klog.Errorf("failed to query user model: %s", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.NewNotFoundError(err)
		}
		return nil, bizerr.NewExternalError(err)
	}
	resp.UserId = userInfo.ID

	res, err := db.GetRedis().Get(context.Background(), getFrozenDurationKey(userInfo.ID)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return resp, nil
		}
		return nil, bizerr.NewExternalError(err)
	}
	// Proactively generating a rpc business error
	return resp, bizerr.NewAuthenticationError(getFrozenError(res))
}

func getFrozenDurationKey(id int32) string {
	return fmt.Sprintf("account_%d_frozen_duration", id)
}

func getFrozenError(frozenDuration string) error {
	return errors.New(fmt.Sprintf("account frozen, please try again in %s minutes", frozenDuration))
}
