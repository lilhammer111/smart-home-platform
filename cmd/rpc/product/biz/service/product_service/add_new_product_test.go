package product_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/binding"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"testing"
)

func TestAddNewProduct_Run(t *testing.T) {
	binding.SetConfPath("", "../../../conf/dev_conf.yaml")
	binding.SetEnvPath("../../../conf/.env")
	ctx := context.Background()
	s := NewAddNewProductService(ctx)
	// init req and assert value

	req := &product.AddProductReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
