package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
	"testing"
)

func TestGetUsersByFilter_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUsersByFilterService(ctx)
	// init req and assert value

	req := &user.UsersFilterReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
