package device_status

import (
	"context"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	device_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device_status"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetDeviceStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetDeviceStatusService(Context context.Context, RequestContext *app.RequestContext) *GetDeviceStatusService {
	return &GetDeviceStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *GetDeviceStatusService) Do(req *common_http.Req) (resp *device_status.DeviceStatusInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
