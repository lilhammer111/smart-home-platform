package brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type GetBrandListService struct {
	ctx context.Context
}

// NewGetBrandListService new GetBrandListService
func NewGetBrandListService(ctx context.Context) *GetBrandListService {
	return &GetBrandListService{ctx: ctx}
}

// Run create note info
func (s *GetBrandListService) Run(req *common.PageFilter) (resp []*product.BrandListResp, err error) {
	brandInfos := make([]model.Brand, 0)
	res := db.GetMysql().Model(&model.Brand{}).Scopes(scope.Paginate(req.Page, req.Limit)).Order("id desc").Find(&brandInfos)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info(utils.InfoNotFound)
		return nil, kerrors.NewBizStatusError(code.NotFound, "No brands.")
	}

	resp = make([]*product.BrandListResp, 0)

	for _, brand := range brandInfos {
		brandListResp := &product.BrandListResp{}
		err = copier.Copy(brandListResp, brand)
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
		}

		err = db.GetMysql().Table(model.Category{}.TableName()).Joins("join category_brands on categories.id = category_brands.category_id").
			Where("category_brands.brand_id = ?", brand.Id).Select("name").Find(&brandListResp.CategoryList).Error
		if err != nil {
			klog.Error(err)
			return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
		}

		resp = append(resp, brandListResp)
	}

	return resp, nil
}
