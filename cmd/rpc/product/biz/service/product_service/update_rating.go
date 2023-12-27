package product_service

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
	"gorm.io/gorm/clause"
)

type UpdateRatingService struct {
	ctx context.Context
}

// NewUpdateRatingService new UpdateRatingService
func NewUpdateRatingService(ctx context.Context) *UpdateRatingService {
	return &UpdateRatingService{ctx: ctx}
}

// Run create note info
func (s *UpdateRatingService) Run(req *product.RatingReq) (resp *product.RatingInfo, err error) {
	tx := db.GetMysql().Begin()
	rating := &model.Rating{}
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("product_id = ?", req.ProductId).First(&rating).Error
	if err != nil {
		klog.Error(err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The product rating is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	// todo test
	res := tx.Model(&rating).
		Updates(map[string]interface{}{
			"cur_rating":     gorm.Expr("(cur_rating + ?) / (total_customer + 1)", req.Rating),
			"total_customer": gorm.Expr("total_customer + 1"),
		})
	if res.Error != nil {
		klog.Error(res.Error)
		tx.Rollback()
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no updates")
		tx.Rollback()
		return nil, kerrors.NewBizStatusError(code.BadRequest, "No updates.")
	}
	tx.Commit()
	resp = &product.RatingInfo{}
	err = copier.Copy(resp, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return resp, nil
}
