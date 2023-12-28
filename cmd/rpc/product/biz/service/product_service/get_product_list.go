package product_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type GetProductListService struct {
	ctx context.Context
}

// NewGetProductListService new GetProductListService
func NewGetProductListService(ctx context.Context) *GetProductListService {
	return &GetProductListService{ctx: ctx}
}

// Run create note info
func (s *GetProductListService) Run(req *product.ProductFilter) (resp []*product.ProductBasicInfo, err error) {
	tx := db.GetMysql()

	// by the specific category
	if req.CategoryId != nil {
		tx = tx.Where("category_id = ?", req.CategoryId)
	}

	// by the specific brands
	if req.BrandIdList != nil {
		tx = tx.Where("brand_id in ?", req.BrandIdList)
	}

	// by product price order
	if req.IsPriceAsc != nil {
		if *req.IsPriceAsc {
			tx = tx.Order("price")
		} else {
			tx = tx.Order("price desc")
		}
	}

	// by product rating order
	if req.IsRatingAsc != nil {
		if *req.IsRatingAsc {
			tx = tx.Order("rating")
		} else {
			tx = tx.Order("rating desc")
		}
	}

	// by product state
	productState := parseStateToInt(req.OnSale, req.IsNew, req.IsHot, req.IsFreeShipping, req.IsRecommended)
	tx = tx.Where("state = ?", productState)

	if req.Search != nil {
		tx = tx.Where("search like %?%", "%"+*req.Search+"%")
	}

	// by pagination
	productInfos := make([]model.Product, 0)
	res := tx.Scopes(scope.Paginate(req.Page, req.Limit)).Find(&productInfos)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info(utils.InfoNotFound)
		return nil, kerrors.NewBizStatusError(code.NotFound, "No eligible products.")
	}

	resp = make([]*product.ProductBasicInfo, 0)
	err = copier.Copy(&resp, productInfos)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}

func parseStateToInt(onSale, isNew, isHot, isFreeShipping, isRecommended *bool) (state model.ProductState) {
	if onSale != nil && *onSale {
		state |= model.OnSale
	}

	if isNew != nil && *isNew {
		state |= model.IsNew
	}

	if isHot != nil && *isHot {
		state |= model.IsHot
	}

	if isFreeShipping != nil && *isFreeShipping {
		state |= model.IsFreeShipping
	}

	if isRecommended != nil && *isRecommended {
		state |= model.IsRecommended
	}
	return
}
