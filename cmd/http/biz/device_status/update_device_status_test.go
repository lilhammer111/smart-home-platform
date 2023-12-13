package device_status

import (
	"context"
	"testing"

	device_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device_status"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUpdateDeviceStatusService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUpdateDeviceStatusService(ctx, c)
	// init req and assert value
	req := &device_status.DeviceStatusInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
