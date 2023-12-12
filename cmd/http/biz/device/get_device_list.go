package device

import (
	"context"

	device "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetDeviceListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetDeviceListService(Context context.Context, RequestContext *app.RequestContext) *GetDeviceListService {
	return &GetDeviceListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetDeviceListService) Do(req *device.DeviceFilter) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
