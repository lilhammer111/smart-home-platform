package model_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"testing"
)

func TestGetAllModels_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetAllModelsService(ctx)
	// init req and assert value
	resp, err := s.Run()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
