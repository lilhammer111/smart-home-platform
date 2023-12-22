package user

import (
	"context"
	"testing"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	user "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetCurUserInfoService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetCurUserInfoService(ctx, c)
	// init req and assert value
	req := &common.Empty{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
