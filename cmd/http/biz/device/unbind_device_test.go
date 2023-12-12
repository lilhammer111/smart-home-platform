package device

import (
	"context"
	"testing"

	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUnbindDeviceService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUnbindDeviceService(ctx, c)
	// init req and assert value
	req := &int32{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
