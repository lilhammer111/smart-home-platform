package product_service

import (
	"context"
	"errors"
	"fmt"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RateProductService struct {
	ctx context.Context
}

// NewRateProductService new RateProductService
func NewRateProductService(ctx context.Context) *RateProductService {
	return &RateProductService{ctx: ctx}
}

// Run create note info
func (s *RateProductService) Run(req *product.RatingReq) (resp *product.RatingResp, err error) {
	tx := db.GetMysql().Begin()
	productRatingInfo := model.ProductRating{}

	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&productRatingInfo, req.ProductId).Error
	if err != nil {
		klog.Error(err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The product is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	res := tx.Model(&productRatingInfo).Updates(map[string]interface{}{
		"total_customer": gorm.Expr("total_customer + ?", 1),
		"total_rating":   gorm.Expr("total_rating + ?", req.Rating),
	})
	if res.Error != nil || res.RowsAffected == 0 {
		klog.Error(res.Error)
		tx.Rollback()
		// If plussing the total customer number fails, response an internal error here.
		return nil, kerrors.NewBizStatusError(code.BadRequest, msg.InternalError)
	}

	err = tx.First(&productRatingInfo, req.ProductId).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	productInfo := model.Product{}
	productInfo.ID = req.ProductId
	res = tx.Model(&productInfo).Select("rating").
		Update("rating", gorm.Expr("? / ?", productRatingInfo.TotalRating, productRatingInfo.TotalCustomer))
	if res.Error != nil {
		klog.Error(res.Error)
		tx.Rollback()
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	// Don't judge the RowsAffected, because the rating is possible to be the same with the previous.
	tx.Commit()

	err = db.GetMysql().First(&productInfo, req.ProductId).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.RatingResp{
		Rating: fmt.Sprintf("%.1f", productInfo.Rating),
	}

	return resp, nil
}
