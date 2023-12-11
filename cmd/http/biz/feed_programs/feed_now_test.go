package feed_programs

import (
	"context"
	"testing"

	feed_programs "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_programs"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestFeedNowService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewFeedNowService(ctx, c)
	// init req and assert value
	req := &feed_programs.FeedNowReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
