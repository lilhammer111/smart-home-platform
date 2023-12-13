package feed_program

import (
	"context"
	"testing"

	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUpdateProgramInfoService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUpdateProgramInfoService(ctx, c)
	// init req and assert value
	req := &feed_program.ProgramInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
