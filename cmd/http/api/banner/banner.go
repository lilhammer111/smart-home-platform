package banner

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"git.zqbjj.top/pet/services/cmd/http/api/banner/handler"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
)

// GetAllBanners .
// @id			GetAllBanners
// @Summary		get all banners
// @Tags		banners
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Success		200				{object}		example.RespOk{data=[]example.BannerData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/banners/list [GET]
func GetAllBanners(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetAllBannersService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddNewBanner .
// @id			AddNewBanner
// @Summary		bind banner info
// @Tags		banners
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		banner	body	example.NewBannerBody	true	"banner data"
// @Success		201				{object}		example.RespCreated{data=example.BannerData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/banners/add [POST]
func AddNewBanner(ctx context.Context, c *app.RequestContext) {
	var err error
	var req banner.NewBanner
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewAddNewBannerService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateBanner .
// @id			UpdateBanner
// @Summary		update banner info
// @Tags		banners
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		banner	body	example.BannerData	true	"banner data"
// @Success		200				{object}		example.RespOk{data=example.BannerData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/banners/update [PUT]
func UpdateBanner(ctx context.Context, c *app.RequestContext) {
	var err error
	var req banner.BannerInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateBannerService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteBanner .
// @id			DeleteBanner
// @Summary		delete a product banner
// @Tags		banners
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"banner id"
// @Success		200				{object}		example.RespOk{data=example.Empty} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/banners/delete [DELETE]
func DeleteBanner(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteBannerService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
