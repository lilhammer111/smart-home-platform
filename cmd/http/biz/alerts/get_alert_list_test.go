package alerts

import (
	"context"
	"testing"

	alerts "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alerts"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetAlertListService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetAlertListService(ctx, c)
	// init req and assert value
	req := &alerts.AlertsFilter{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
