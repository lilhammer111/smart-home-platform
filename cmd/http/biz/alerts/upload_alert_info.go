package alerts

import (
	"context"

	alerts "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alerts"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type UploadAlertInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadAlertInfoService(Context context.Context, RequestContext *app.RequestContext) *UploadAlertInfoService {
	return &UploadAlertInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadAlertInfoService) Do(req *alerts.AlertInfo) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
