package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"testing"
)

func TestQueryUsersWithFilter_Run(t *testing.T) {
	ctx := context.Background()
	s := NewQueryUsersWithFilterService(ctx)
	// init req and assert value

	req := &user.UsersFilter{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}