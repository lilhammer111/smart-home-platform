package auth

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMobileLoginService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewMobileLoginService(ctx, c)
	// init req and assert value
	req := &auth.MobileLoginReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
