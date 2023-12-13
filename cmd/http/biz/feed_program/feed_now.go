package feed_program

import (
	"context"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	"github.com/cloudwego/hertz/pkg/app"
)

type FeedNowService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFeedNowService(Context context.Context, RequestContext *app.RequestContext) *FeedNowService {
	return &FeedNowService{RequestContext: RequestContext, Context: Context}
}

func (h *FeedNowService) Do(req *feed_program.FeedNowReq) (resp *common_http.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
