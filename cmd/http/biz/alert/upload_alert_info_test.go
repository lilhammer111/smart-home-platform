package alert

import (
	"context"
	"testing"

	alert "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUploadAlertInfoService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUploadAlertInfoService(ctx, c)
	// init req and assert value
	req := &alert.AlertInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
