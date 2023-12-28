package handler

import (
	"context"
	"testing"

	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestRateProductService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewRateProductService(ctx, c)
	// init req and assert value
	req := &product.RatingReq{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
