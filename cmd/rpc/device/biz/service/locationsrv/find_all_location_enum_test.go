package locationsrv

import (
	"context"
	"testing"
)

func TestFindAllLocationEnum_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFindAllLocationEnumService(ctx)
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
