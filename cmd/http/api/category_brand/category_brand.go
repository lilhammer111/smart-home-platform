package category_brand

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/category_brand/handler"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// BatchAddCategoryBrand .
// @id			BatchAddCategoryBrand
// @Summary		add related categories in batch by brand id
// @Tags		product category_brand
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		category_brand	body	example.NewCategoryBrandBody	true	"category_brand data"
// @Success		201				{object}		example.RespCreated{data=example.CategoryBrandData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/category_brand/batch_add [POST]
func BatchAddCategoryBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category_brand.NewCategoryBrand
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
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

// BatchReduceCategoryBrand .
// @id			BatchReduceCategoryBrand
// @Summary		reduce related categories in batch by brand id
// @Tags		product category_brand
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		category_brand	body	example.NewCategoryBrandBody	true	"category data"
// @Success		201				{object}		example.RespNoContent				 "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/category_brand/batch_reduce [DELETE]
func BatchReduceCategoryBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category_brand.NewCategoryBrand
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewBatchReduceCategoryBrandService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteBrandByCategory .
// @id			DeleteBrandByCategory
// @Summary		delete related brand by category id
// @Tags		product category_brand
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"id"
// @Success		201				{object}		example.RespNoContent				 "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/category_brand/delete_brand [DELETE]
func DeleteBrandByCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteBrandByCategoryService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteCategoryByBrand .
// @id			DeleteCategoryByBrand
// @Summary		delete related category by brand id
// @Tags		product category_brand
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"id"
// @Success		201				{object}		example.RespNoContent				 "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/category_brand/delete_category [DELETE]
func DeleteCategoryByBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
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
