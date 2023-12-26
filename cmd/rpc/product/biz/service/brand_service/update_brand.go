package brand_service

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

type UpdateBrandService struct {
	ctx context.Context
}

// NewUpdateBrandService new UpdateBrandService
func NewUpdateBrandService(ctx context.Context) *UpdateBrandService {
	return &UpdateBrandService{ctx: ctx}
}

// Run create note info
func (s *UpdateBrandService) Run(req *product.BrandInfo) (resp *product.BrandInfo, err error) {
	err = db.GetMysql().First(&model.Brand{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The brand you'd like to update is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	res := db.GetMysql().Model(&model.Brand{Id: req.Id}).Updates(req)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no updates")
		return nil, kerrors.NewBizStatusError(code.BadRequest, "No updates.")
	}

	resp = &product.BrandInfo{}
	err = copier.Copy(resp, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return resp, nil
}
