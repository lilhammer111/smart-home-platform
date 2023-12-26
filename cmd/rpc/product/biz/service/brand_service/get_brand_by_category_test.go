package brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"testing"
)

func TestGetBrandByCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetBrandByCategoryService(ctx)
	// init req and assert value

	req := &product.GetBrandByCatReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
