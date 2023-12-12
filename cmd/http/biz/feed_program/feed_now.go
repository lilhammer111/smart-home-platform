package feed_program

import (
	"context"

	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type FeedNowService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFeedNowService(Context context.Context, RequestContext *app.RequestContext) *FeedNowService {
	return &FeedNowService{RequestContext: RequestContext, Context: Context}
}

func (h *FeedNowService) Do(req *feed_program.FeedNowReq) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
