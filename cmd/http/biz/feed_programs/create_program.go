package feed_programs

import (
	"context"

	feed_programs "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_programs"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProgramService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProgramService(Context context.Context, RequestContext *app.RequestContext) *CreateProgramService {
	return &CreateProgramService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProgramService) Do(req *feed_programs.ProgramInfo) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
