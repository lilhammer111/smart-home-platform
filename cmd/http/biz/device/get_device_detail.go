package device

import (
	"context"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	device "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetDeviceDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetDeviceDetailService(Context context.Context, RequestContext *app.RequestContext) *GetDeviceDetailService {
	return &GetDeviceDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetDeviceDetailService) Do(req *common_http.Req) (resp *device.DeviceInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
