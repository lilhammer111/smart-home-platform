package auth

import (
	"context"
	"testing"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestPwdLoginService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewPwdLoginService(ctx, c)
	// init req and assert value
	req := &auth.PwdLoginReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
