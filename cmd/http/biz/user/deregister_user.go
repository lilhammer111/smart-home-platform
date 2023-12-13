package user

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/rpc_client/user_micro"

	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeregisterUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeregisterUserService(Context context.Context, RequestContext *app.RequestContext) *DeregisterUserService {
	return &DeregisterUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeregisterUserService) Do(req *standard.Req) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcReq := &common.IdRpcReq{Id: req.Id}
	_, err = user_micro.DefaultClient().DeleteUser(h.Context, rpcReq)
	if err != nil {
		resp = &standard.Resp{
			Success: true,
			Code:    204,
			Message: "delete success",
		}
	}

	return
}
