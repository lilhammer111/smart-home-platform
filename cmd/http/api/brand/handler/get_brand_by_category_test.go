package handler

import (
	"context"
	"testing"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetBrandByCategoryService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetBrandByCategoryService(ctx, c)
	// init req and assert value
	req := &brand.GetBrandByCatReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
