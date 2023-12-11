package service

import (
	"context"
	"testing"
)

func TestFindUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFindUserService(ctx)
	// init req and assert value

	req := &int32{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
