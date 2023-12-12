package alert

import (
	"context"
	"testing"

	alert "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetAlertDetailService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetAlertDetailService(ctx, c)
	// init req and assert value
	req := &int32{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
