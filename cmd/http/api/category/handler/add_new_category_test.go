package handler

import (
	"context"
	"testing"

	category "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestAddNewCategoryService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewAddNewCategoryService(ctx, c)
	// init req and assert value
	req := &category.AddCategoryReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
