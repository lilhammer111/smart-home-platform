package service

import (
	"context"
	"fmt"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

const (
	initialFrozenDuration = 5 * time.Minute
)

type FreezePatrolAfterVerifyService struct {
	ctx context.Context
}

// NewFreezePatrolAfterVerifyService new FreezePatrolAfterVerifyService
func NewFreezePatrolAfterVerifyService(ctx context.Context) *FreezePatrolAfterVerifyService {
	return &FreezePatrolAfterVerifyService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolAfterVerifyService) Run(req *micro_user.RpcAfterVerifyReq) (resp *common.Empty, err error) {
	if req.VerifyPassed {
		err = db.GetRedis().Set(context.Background(), getVerifyFailedKey(req.UserId), 0, 0).Err()
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
		}

		return &common.Empty{}, nil
	}
	// If verification failed, let the total number of failures be added by one
	count, err := db.GetRedis().Incr(context.Background(), getVerifyFailedKey(req.UserId)).Result()
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	// If it's first time failing verification
	var currentDuration, lastDuration string
	if count == 1 {
		err = db.GetRedis().SetEx(context.Background(), getFrozenDurationKey(req.UserId), initialFrozenDuration, initialFrozenDuration).Err()
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
		}
		currentDuration = time.Now().Add(initialFrozenDuration).String()
	} else {
		lastDuration, err = db.GetRedis().Get(context.Background(), getFrozenDurationKey(req.UserId)).Result()
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
		}
		var duration int64
		duration, err = strconv.ParseInt(lastDuration, 10, 0)
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
		}

		err = db.GetRedis().SetEx(context.Background(), getFrozenDurationKey(req.UserId), duration*2, time.Duration(duration*2)).Err()
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
		}
		currentDuration = time.Now().Add(time.Duration(duration * 2)).String()
	}

	return &common.Empty{}, kerrors.NewBizStatusError(code.AuthenticationFailed, fmt.Sprintf("account frozen, please try again in %s minutes", currentDuration))
}

func getVerifyFailedKey(id int32) string {
	return fmt.Sprintf("account_%d_failed_count", id)
}
