package feed_program

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/feed_program"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetSelfProgramList .
// @router /api/programs/list [GET]
func GetSelfProgramList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req standard.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetSelfProgramListService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetProgramDetail .
// @router /api/programs/detail [GET]
func GetProgramDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req standard.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetProgramDetailService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProgramInfo .
// @router /api/programs/update [PUT]
func UpdateProgramInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed_program.ProgramInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUpdateProgramInfoService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateProgram .
// @router /api/programs/create [POST]
func CreateProgram(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed_program.ProgramInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewCreateProgramService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// FeedNow .
// @router /api/programs/feed_now [POST]
func FeedNow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed_program.FeedNowReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewFeedNowService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteProgram .
// @router /api/programs/delete [DELETE]
func DeleteProgram(ctx context.Context, c *app.RequestContext) {
	var err error
	var req standard.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewDeleteProgramService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
