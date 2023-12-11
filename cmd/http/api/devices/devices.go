package devices

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/devices"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetDeviceList .
// @router /api/devices [GET]
func GetDeviceList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req devices.DevicesFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetDeviceListService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetDeviceDetail .
// @router /api/devices/detail [GET]
func GetDeviceDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetDeviceDetailService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateDeviceInfo .
// @router /api/devices [PUT]
func UpdateDeviceInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req devices.DeviceInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUpdateDeviceInfoService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// BindDevice .
// @router /api/devices [POST]
func BindDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req devices.DeviceInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewBindDeviceService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UnbindDevice .
// @router /api/devices [DELETE]
func UnbindDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUnbindDeviceService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
