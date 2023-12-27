package category_brand_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"testing"
)

func TestBatchReduceCategoryBrand_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBatchReduceCategoryBrandService(ctx)
	// init req and assert value

	req := &product.NewCategoryBrand_{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
