package brand

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/brand/handler"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// AddNewBrand .
// @id			AddNewBrand
// @Summary		bind brand info
// @Tags		product brands
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		brand	body	example.NewBrandBody	true	"brand data"
// @Success		201				{object}		example.RespCreated{data=example.DeviceData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/brands/add [POST]
func AddNewBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req brand.NewBrand
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
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
// @id			UpdateBrand
// @Summary		update brand info
// @Tags		product brands
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		brand	body	example.BrandData	true	"brand data"
// @Success		200				{object}		example.RespOk{data=example.BrandData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/brands/update [PUT]
func UpdateBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req brand.BrandInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
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

// GetBrandDetail .
// @id			GetBrandDetail
// @Summary		get brand detail
// @Tags		product brands
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id		query	int	true	"id"
// @Success		200				{object}		example.RespOk{data=example.BrandData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/brands/detail [GET]
func GetBrandDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetBrandDetailService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteBrand .
// @id			DeleteBrand
// @Summary		delete brand record
// @Tags		product brands
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"id"
// @Success		204				{object}		example.RespNoContent				"success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/brands/delete [DELETE]
func DeleteBrand(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteBrandService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetRelatedBrandsByCategoryId .
// @id			GetRelatedBrandsByCategoryId
// @Summary		get brands list by category id
// @Tags		product brands
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		category_id		query	int	true	"category_id"
// @Success		200				{object}		example.RespOk{data=[]example.BrandData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/brands/related [GET]
func GetRelatedBrandsByCategoryId(ctx context.Context, c *app.RequestContext) {
	var err error
	var req brand.BrandByCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetRelatedBrandsByCategoryIdService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetBrandList .
// @id			GetBrandList
// @Summary		get brand list
// @Tags		product brands
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		page	query	string	false	"page"
// @Param		limit	query	string	false	"limit"
// @Success		200				{object}		example.RespOk{data=[]example.BrandData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/brands/all [GET]
func GetBrandList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.PageFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
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
