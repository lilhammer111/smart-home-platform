package feed_program

import (
	"context"
	"testing"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestFeedNowService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewFeedNowService(ctx, c)
	// init req and assert value
	req := &feed_program.FeedNowReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
