package handler

import (
	"context"
	"testing"

	banner "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUpdateBannerService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUpdateBannerService(ctx, c)
	// init req and assert value
	req := &banner.BannerInfo{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
