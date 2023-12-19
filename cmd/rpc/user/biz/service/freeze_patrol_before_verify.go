package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type FreezePatrolBeforeVerifyService struct {
	ctx context.Context
} // NewFreezePatrolBeforeVerifyService new FreezePatrolBeforeVerifyService
func NewFreezePatrolBeforeVerifyService(ctx context.Context) *FreezePatrolBeforeVerifyService {
	return &FreezePatrolBeforeVerifyService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolBeforeVerifyService) Run(req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	if err = db.GetMysql().Select()
	return
}
