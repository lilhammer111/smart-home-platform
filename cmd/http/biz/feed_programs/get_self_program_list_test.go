package feed_programs

import (
	"context"
	"testing"

	req "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetSelfProgramListService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetSelfProgramListService(ctx, c)
	// init req and assert value
	req := &req.IdReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
