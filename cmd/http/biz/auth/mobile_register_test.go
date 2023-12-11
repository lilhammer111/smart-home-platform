package auth

import (
	"context"
	"testing"

	auth "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMobileRegisterService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewMobileRegisterService(ctx, c)
	// init req and assert value
	req := &auth.MobileRegisterReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
