package brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"testing"
)

func TestGetBrandList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetBrandListService(ctx)
	// init req and assert value

	req := &product.PageFilter{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
