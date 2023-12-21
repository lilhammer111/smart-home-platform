package user

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
)

type GetUserDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserDetailService(Context context.Context, RequestContext *app.RequestContext) *GetUserDetailService {
	return &GetUserDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserDetailService) Do(req *common.Req) (resp *user.UserInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userInfo, err := micro_user_cli.FindUser(h.Context, &micro_user.RpcFindUserReq{UserId: req.Id})
	if err != nil {
		return nil, err
	}
	resp = &user.UserInfo{}
	if err = copier.Copy(resp, userInfo); err != nil {
		return nil, err
	}
	return
}
