package category_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type AddNewCategoryService struct {
	ctx context.Context
}

// NewAddNewCategoryService new AddNewCategoryService
func NewAddNewCategoryService(ctx context.Context) *AddNewCategoryService {
	return &AddNewCategoryService{ctx: ctx}
}

// Run create note info
func (s *AddNewCategoryService) Run(req *product.NewCategory_) (resp *product.CategoryInfo, err error) {
	categoryInfo := model.Category{}
	err = copier.Copy(&categoryInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&categoryInfo).Error
	if err != nil {
		klog.Error(err)
		if isConflict, kerr := utils.CheckResourceConflict(err, "Category name conflicts."); isConflict {
			return nil, kerr
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.CategoryInfo{}
	err = copier.Copy(resp, categoryInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
