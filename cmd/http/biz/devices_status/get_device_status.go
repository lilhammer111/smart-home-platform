package devices_status

import (
	"context"

	req "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetDeviceStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetDeviceStatusService(Context context.Context, RequestContext *app.RequestContext) *GetDeviceStatusService {
	return &GetDeviceStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *GetDeviceStatusService) Do(req *req.IdReq) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
