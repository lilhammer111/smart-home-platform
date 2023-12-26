package category_brand

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"git.zqbjj.top/pet/services/cmd/http/api/category_brand/handler"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
)

// GetCategoryBrandList .
// @router /api/products/category_brand/list [GET]
func GetCategoryBrandList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetCategoryBrandListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// BatchAddCategoryBrand .
// @router /api/products/category_brand/batch_add [POST]
func BatchAddCategoryBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category_brand.NewCategoryBrand
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewBatchAddCategoryBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateCategoryBrand .
// @router /api/products/category_brand/update [PUT]
func UpdateCategoryBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category_brand.CategoryBrandInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateCategoryBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteCategoryByBrand .
// @router /api/products/category_brand/delete_category [DELETE]
func DeleteCategoryByBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteCategoryByBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
