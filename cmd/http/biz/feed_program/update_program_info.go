package feed_program

import (
	"context"

	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProgramInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProgramInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateProgramInfoService {
	return &UpdateProgramInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProgramInfoService) Do(req *feed_program.ProgramInfo) (resp *feed_program.ProgramInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
