package devices

import (
	"context"
	"testing"

	devices "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUpdateDeviceInfoService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUpdateDeviceInfoService(ctx, c)
	// init req and assert value
	req := &devices.DeviceInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
