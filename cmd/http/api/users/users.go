package users

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/users"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/users"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetUserList .
// @router /api/users [GET]
func GetUserList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.UsersFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetUserListService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetUserDetail .
// @router /api/users/detail [GET]
func GetUserDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
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
// @router /api/users [PUT]
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
// @router /api/users [DELETE]
func DeregisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
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
