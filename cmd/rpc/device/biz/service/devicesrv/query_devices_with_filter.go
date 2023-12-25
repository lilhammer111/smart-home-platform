package devicesrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type QueryDevicesWithFilterService struct {
	ctx context.Context
}

// NewQueryDevicesWithFilterService new QueryDevicesWithFilterService
func NewQueryDevicesWithFilterService(ctx context.Context) *QueryDevicesWithFilterService {
	return &QueryDevicesWithFilterService{ctx: ctx}
}

// Run create note info
func (s *QueryDevicesWithFilterService) Run(req *device.DeviceFilter) (resp []*device.DeviceInfo, err error) {
	tx := db.GetMysql().Model(&model.Device{})

	if req.State != nil {
		tx = tx.Where("state = ?", req.State)
	}

	if req.Sorts != nil {
		for _, sortCond := range req.Sorts {
			tx = tx.Order(sortCond)
		}
	}

	if req.StartDate != nil && req.EndDate != nil {
		tx = tx.Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	} else {
		oneMonthAgo := time.Now().AddDate(0, -1, 0)
		tx = tx.Where("created_at >= ?", oneMonthAgo)
	}

	resp = make([]*device.DeviceInfo, 0)
	if req.Page != nil && req.Limit != nil {
		err = tx.Scopes(scope.Paginate(req.Page, req.Limit)).Find(&resp).Error
	} else {
		err = tx.Offset(0).Limit(10).Find(&resp).Error
	}
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	return resp, nil
}
