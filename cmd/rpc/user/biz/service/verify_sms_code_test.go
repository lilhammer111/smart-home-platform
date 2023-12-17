package service

import (
	"context"
	"testing"
)

func TestVerifySmsCode_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifySmsCodeService(ctx)
	// init req and assert value

	mobile := &string{}

	smsCode := &string{}
	resp, err := s.Run(mobilesmsCode)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
