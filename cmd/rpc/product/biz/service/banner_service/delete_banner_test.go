package banner_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"testing"
)

func TestDeleteBanner_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteBannerService(ctx)
	// init req and assert value

	req := &common.Req{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
