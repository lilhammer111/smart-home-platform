package devices_status

import (
	"context"
	"testing"

	devices_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices_status"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestInitDeviceStatusService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewInitDeviceStatusService(ctx, c)
	// init req and assert value
	req := &devices_status.DeviceStatusInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
