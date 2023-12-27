package category_service

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

type UpdateCategoryService struct {
	ctx context.Context
}

// NewUpdateCategoryService new UpdateCategoryService
func NewUpdateCategoryService(ctx context.Context) *UpdateCategoryService {
	return &UpdateCategoryService{ctx: ctx}
}

// Run create note info
func (s *UpdateCategoryService) Run(req *product.CategoryInfo) (resp *product.CategoryInfo, err error) {
	err = db.GetMysql().First(&model.Category{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The category you'd like to update is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	categoryInfo := model.Category{}
	err = copier.Copy(&categoryInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	res := db.GetMysql().Updates(&categoryInfo)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no updates")
		return nil, kerrors.NewBizStatusError(code.BadRequest, "No updates.")
	}

	resp = &product.CategoryInfo{}
	err = copier.Copy(resp, categoryInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return resp, nil
}
