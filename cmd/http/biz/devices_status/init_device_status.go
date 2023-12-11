package devices_status

import (
	"context"

	devices_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices_status"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type InitDeviceStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewInitDeviceStatusService(Context context.Context, RequestContext *app.RequestContext) *InitDeviceStatusService {
	return &InitDeviceStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *InitDeviceStatusService) Do(req *devices_status.DeviceStatusInfo) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
