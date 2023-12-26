package brand

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"git.zqbjj.top/pet/services/cmd/http/api/brand/handler"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
)

// GetBrandList .
// @router /api/products/brands/list [GET]
func GetBrandList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.PageFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetBrandListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetBrandByCategory .
// @router /api/products/categories/brands/list [GET]
func GetBrandByCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req brand.GetBrandByCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetBrandByCategoryService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddNewBrand .
// @router /api/products/brands/add [POST]
func AddNewBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req brand.NewBrand
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewAddNewBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateBrand .
// @router /api/products/brands/update [PUT]
func UpdateBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req brand.BrandInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DelteBrand .
// @router /api/products/brands/delete [DELETE]
func DelteBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDelteBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
