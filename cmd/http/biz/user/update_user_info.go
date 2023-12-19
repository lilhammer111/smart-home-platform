package user

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"
	"github.com/jinzhu/copier"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	rpcUser "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Do(req *user.UserInfo) (resp *user.UserInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	rpcReq := rpcUser.UserInfo{}
	if err = copier.Copy(&rpcReq, req); err != nil {
		return nil, err
	}
	userInfo, err := micro_user_cli.UpdateUser(h.Context, &rpcReq)
	if err != nil {
		return nil, err
	}
	if err = copier.Copy(resp, userInfo); err != nil {
		return nil, err
	}
	return
}
