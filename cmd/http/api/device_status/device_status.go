package device_status

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device_status"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/device_status"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetDeviceStatus .
// @router /api/devices/status/detail [GET]
func GetDeviceStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req standard.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetDeviceStatusService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// InitDeviceStatus .
// @router /api/devices/status/init [POST]
func InitDeviceStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device_status.DeviceStatusInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewInitDeviceStatusService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateDeviceStatus .
// @router /api/devices/status/update [PUT]
func UpdateDeviceStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device_status.DeviceStatusInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUpdateDeviceStatusService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
