package device_status

import (
	"context"
	"testing"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	device_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device_status"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetDeviceStatusService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetDeviceStatusService(ctx, c)
	// init req and assert value
	req := &common_http.Req{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
