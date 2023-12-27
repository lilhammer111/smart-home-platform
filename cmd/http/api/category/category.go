package category

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/category/handler"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// GetCategoryList .
// @id			GetCategoryList
// @Summary		get category list
// @Tags		product categories
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		page	query	string	false	"page"
// @Param		limit	query	string	false	"limit"
// @Success		200				{object}		example.RespOk{data=[]example.CategoryData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/categories/list [GET]
func GetCategoryList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.PageFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetCategoryListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCategoryDetail .
// @id			GetCategoryDetail
// @Summary		get category detail
// @Tags		product categories
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id		query	int	true	"id"
// @Success		200				{object}		example.RespOk{data=example.CategoryData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/categories/detail [GET]
func GetCategoryDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetCategoryDetailService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddNewCategory .
// @id			AddNewCategory
// @Summary		add a new category
// @Tags		product categories
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		category	body	example.NewCategoryBody	true	"category data"
// @Success		201				{object}		example.RespCreated{data=example.CategoryData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/categories/add [POST]
func AddNewCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.NewCategory
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewAddNewCategoryService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateCategory .
// @id			UpdateCategory
// @Summary		update category info
// @Tags		product categories
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		category	body	example.CategoryData	true	"category data"
// @Success		200				{object}		example.RespOk{data=example.CategoryData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/categories/update [PUT]
func UpdateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateCategoryService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteCategory .
// @id			DeleteCategory
// @Summary		delete category record
// @Tags		product categories
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"id"
// @Success		204				{object}		example.RespNoContent				"success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/categories/delete [DELETE]
func DeleteCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteCategoryService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
