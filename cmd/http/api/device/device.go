package device

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/device"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetDeviceList .
// @router /api/devices/list [GET]
func GetDeviceList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.DeviceFilter
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
	var req standard.Req
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
// @router /api/devices/update [PUT]
func UpdateDeviceInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.DeviceInfo
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
// @router /api/devices/bind [POST]
func BindDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.DeviceInfo
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
// @router /api/devices/unbind [DELETE]
func UnbindDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req standard.Req
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
