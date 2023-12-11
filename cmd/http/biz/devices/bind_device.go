package devices

import (
	"context"

	devices "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type BindDeviceService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBindDeviceService(Context context.Context, RequestContext *app.RequestContext) *BindDeviceService {
	return &BindDeviceService{RequestContext: RequestContext, Context: Context}
}

func (h *BindDeviceService) Do(req *devices.DeviceInfo) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
