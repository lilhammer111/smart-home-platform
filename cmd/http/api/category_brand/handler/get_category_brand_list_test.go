package handler

import (
	"context"
	"testing"

	category_brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetCategoryBrandListService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetCategoryBrandListService(ctx, c)
	// init req and assert value
	req := &common.Req{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
