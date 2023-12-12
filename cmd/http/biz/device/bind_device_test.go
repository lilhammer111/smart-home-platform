package device

import (
	"context"
	"testing"

	device "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestBindDeviceService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewBindDeviceService(ctx, c)
	// init req and assert value
	req := &device.DeviceInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
