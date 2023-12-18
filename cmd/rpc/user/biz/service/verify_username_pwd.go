package service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type VerifyUsernamePwdService struct {
	ctx context.Context
} // NewVerifyUsernamePwdService new VerifyUsernamePwdService
func NewVerifyUsernamePwdService(ctx context.Context) *VerifyUsernamePwdService {
	return &VerifyUsernamePwdService{ctx: ctx}
}

// Run create note info
func (s *VerifyUsernamePwdService) Run(req *micro_user.RpcVerifyUsernamePwdReq) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
