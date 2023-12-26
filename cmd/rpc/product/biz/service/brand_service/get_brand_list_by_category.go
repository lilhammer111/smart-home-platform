package brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetBrandListByCategoryService struct {
	ctx context.Context
}

// NewGetBrandListByCategoryService new GetBrandListByCategoryService
func NewGetBrandListByCategoryService(ctx context.Context) *GetBrandListByCategoryService {
	return &GetBrandListByCategoryService{ctx: ctx}
}

// Run create note info
func (s *GetBrandListByCategoryService) Run(req *product.BrandByCatReq) (resp []*product.BrandInfo, err error) {
	brandIds := make([]int32, 0)
	res := db.GetMysql().Model(&model.CategoryBrand{Id: req.CategoryId}).Pluck("brand_id", brandIds)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no related brands")
		return nil, kerrors.NewBizStatusError(code.NotFound, "No related brands.")
	}

	res = db.GetMysql().Model(&model.Brand{}).Where("id in ?", brandIds).Find(&resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no related brands")
		return nil, kerrors.NewBizStatusError(code.NotFound, "No related brands.")
	}

	return resp, nil
}
