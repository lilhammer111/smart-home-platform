package device

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/device/handler"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// GetDeviceList .
// @id			GetDeviceList
// @Summary		get device list
// @Tags		devices
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		page	query	string	false	"page"
// @Param		limit	query	string	false	"limit"
// @Param		state	query	string	false	"state"
// @Param		state	query	string	false	"sorts"
// @Param		state	query	string	false	"search"
// @Param		state	query	string	false	"start_date"
// @Param		state	query	string	false	"end_date"
// @Success		200				{object}		example.RespOk{data=[]example.DeviceData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/list [GET]
func GetDeviceList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.DeviceFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetDeviceListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetDeviceDetail .
// @id			GetDeviceDetail
// @Summary		get device detail
// @Tags		devices
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id		query	int	true	"id"
// @Success		200				{object}		example.RespOk{data=example.DeviceData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/detail [GET]
func GetDeviceDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetDeviceDetailService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateDeviceInfo .
// @id			UpdateDeviceInfo
// @Summary		update device info
// @Tags		devices
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		devices	body	example.DeviceData	true	"device data"
// @Success		200				{object}		example.RespOk{data=example.DeviceData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/update [PUT]
func UpdateDeviceInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.DeviceInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateDeviceInfoService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// BindDevice .
// @id			BindDevice
// @Summary		bind device info
// @Tags		devices
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		devices	body	example.DeviceData	true	"device data"
// @Success		200				{object}		example.RespOk{data=example.DeviceData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/bind [POST]
func BindDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.DeviceInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewBindDeviceService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UnbindDevice .
// @id			UnbindDevice
// @Summary		unbind device
// @Tags		devices
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"device id"
// @Success		204				{object}		example.RespNoContent				"success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/unbind [DELETE]
func UnbindDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUnbindDeviceService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
