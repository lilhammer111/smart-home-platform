package alertsrv

import (
	"context"
	alert "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
)

type QueryAlertsWithFilterService struct {
	ctx context.Context
} // NewQueryAlertsWithFilterService new QueryAlertsWithFilterService
func NewQueryAlertsWithFilterService(ctx context.Context) *QueryAlertsWithFilterService {
	return &QueryAlertsWithFilterService{ctx: ctx}
}

// Run create note info
func (s *QueryAlertsWithFilterService) Run(req *alert.AlertFilter) (resp []*alert.AlertInfo, err error) {
	// Finish your business logic.

	return
}
