package banner_service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UpdateBannerService struct {
	ctx context.Context
}

// NewUpdateBannerService new UpdateBannerService
func NewUpdateBannerService(ctx context.Context) *UpdateBannerService {
	return &UpdateBannerService{ctx: ctx}
}

// Run create note info
func (s *UpdateBannerService) Run(req *product.BannerInfo) (resp *product.BannerInfo, err error) {
	err = db.GetMysql().First(&model.Banner{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The banner you'd like to update is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	res := db.GetMysql().Model(&model.Banner{Id: req.Id}).Updates(req)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no updates")
		return nil, kerrors.NewBizStatusError(code.BadRequest, "No updates.")
	}

	resp = &product.BannerInfo{}
	err = copier.Copy(resp, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return resp, nil
}
