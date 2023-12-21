package user

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/user"
)

// GetUserList .
// @Summary		get user list
// @Tags		users
// @Produce		json
// @Param		page	query	string	false	"page"
// @Param		limit	query	string	false	"limit"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/users/list [GET]
func GetUserList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UsersFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetUserListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Set(responder.SuccessMessage, "getting user list success")
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetUserDetail .
// @Summary		get user detail
// @Tags		users
// @Produce		json
// @Param		id		path	int	true	"id"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/users/detail [GET]
func GetUserDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetUserDetailService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Set(responder.SuccessMessage, "getting user details success")
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateUserInfo .
// @Summary		update user info
// @Tags		users
// @Access		json
// @Produce		json
// @Param		users	body	user.UserInfo	true	"user data"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/users/update [PUT]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUpdateUserInfoService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Set(responder.SuccessMessage, "updating user information success")
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeregisterUser .
// @Summary		deregister user
// @Tags		users
// @Produce		json
// @Param		id	query	string	true	"user id"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/users/deregister [DELETE]
func DeregisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewDeregisterUserService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Set(responder.SuccessMessage, "deregister success")
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
