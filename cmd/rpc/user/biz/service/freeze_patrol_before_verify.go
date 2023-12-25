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
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
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
			return nil, kerrors.NewBizStatusError(code.NotFound, "account is no found")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	resp.UserId = userInfo.ID

	// get the ttl
	remainingTime, err := db.GetRedis().TTL(context.Background(), getFrozenDurationKey(userInfo.ID)).Result()
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	if remainingTime == -2 {
		klog.Info("account is in normal")
		return resp, nil
	} else if remainingTime == -1 {
		klog.Info("key do not set the ttl")
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	} else {
		// Proactively generating a rpc business error
		return resp, kerrors.NewBizStatusError(code.AuthenticationFailed, fmt.Sprintf("account frozen, please try again in %s minutes", remainingTime))
	}
}

func getFrozenDurationKey(id int32) string {
	return fmt.Sprintf("account_%d_frozen_duration", id)
}

func getFrozenError(frozenDuration string) error {
	return errors.New(fmt.Sprintf("account frozen, please try again in %s minutes", frozenDuration))
}
