package handler

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/http/mw"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetCurUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCurUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetCurUserInfoService {
	return &GetCurUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCurUserInfoService) Do(req *common.Empty) (resp *user.UserInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	uid := h.RequestContext.GetInt32(mw.IdentityKey)
	if uid == 0 {
		hlog.Info("payload do not exist")
		return nil, errors.New("failed to get id payload")
	}
	hlog.Infof("user id in payload: %d", uid)

	userInfo, err := micro_user_cli.FindUser(h.Context, &micro_user.RpcFindUserReq{UserId: uid})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &user.UserInfoResp{}
	err = copier.Copy(resp, userInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return resp, nil
}
