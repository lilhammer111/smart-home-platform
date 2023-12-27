package brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetRelatedBrandsByCategoryIdService struct {
	ctx context.Context
}

// NewGetRelatedBrandsByCategoryIdService new GetRelatedBrandsByCategoryIdService
func NewGetRelatedBrandsByCategoryIdService(ctx context.Context) *GetRelatedBrandsByCategoryIdService {
	return &GetRelatedBrandsByCategoryIdService{ctx: ctx}
}

// Run create note info
func (s *GetRelatedBrandsByCategoryIdService) Run(req *product.BrandByCatReq) (resp []*product.BrandInfo, err error) {
	brandIds := make([]int32, 0)
	res := db.GetMysql().Model(&model.CategoryBrand{Id: req.CategoryId}).Pluck("brand_id", &brandIds)
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

	// res := db.GetMysql().
	//    Model(&model.CategoryBrand{}).
	//    Select("brands.*").
	//    Joins("join brands on brands.id = category_brands.brand_id").
	//    Where("category_brands.category_id = ?", req.CategoryId).
	//    Find(&resp)

	return resp, nil
}
