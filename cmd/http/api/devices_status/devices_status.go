package devices_status

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices_status"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/devices_status"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetDeviceStatus .
// @router /api/devices/status [GET]
func GetDeviceStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
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
// @router /api/devices/status [POST]
func InitDeviceStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req devices_status.DeviceStatusInfo
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
// @router /api/devices/status [PUT]
func UpdateDeviceStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req devices_status.DeviceStatusInfo
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
