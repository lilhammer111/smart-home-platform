package feed_program

import (
	"context"

	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProgramDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProgramDetailService(Context context.Context, RequestContext *app.RequestContext) *GetProgramDetailService {
	return &GetProgramDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProgramDetailService) Do(req *standard.Req) (resp *feed_program.ProgramInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
