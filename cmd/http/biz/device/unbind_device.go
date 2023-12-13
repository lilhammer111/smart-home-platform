package device

import (
	"context"

	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type UnbindDeviceService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUnbindDeviceService(Context context.Context, RequestContext *app.RequestContext) *UnbindDeviceService {
	return &UnbindDeviceService{RequestContext: RequestContext, Context: Context}
}

func (h *UnbindDeviceService) Do(req *standard.Req) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
