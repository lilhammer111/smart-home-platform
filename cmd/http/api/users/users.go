package users

import (
	"context"

	biz "git.zqbjj.top/pet/services/pet-feeder/cmd/http/biz/users"
	"git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/users"
	"git.zqbjj.top/pet/services/pet-feeder/cmd/http/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetUserDetail .
// @router /api/users/detail [GET]
func GetUserDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.UserID
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetUserDetailService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateUserInfo .
// @router /api/users/detail [PUT]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.UserInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUpdateUserInfoService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeregisterUser .
// @router /api/users/detail [GET]
func DeregisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.UserID
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewDeregisterUserService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
