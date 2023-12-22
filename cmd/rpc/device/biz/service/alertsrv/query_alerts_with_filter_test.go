package alertsrv

import (
	"context"
	alert "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
	"testing"
)

func TestQueryAlertsWithFilter_Run(t *testing.T) {
	ctx := context.Background()
	s := NewQueryAlertsWithFilterService(ctx)
	// init req and assert value

	req := &alert.AlertFilter{
		Page:      nil,
		Limit:     nil,
		Level:     nil,
		DeviceId:  nil,
		Sorts:     nil,
		StartDate: nil,
		EndDate:   nil,
		IsOngoing: nil,
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
