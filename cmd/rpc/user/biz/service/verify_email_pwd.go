package service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type VerifyEmailPwdService struct {
	ctx context.Context
} // NewVerifyEmailPwdService new VerifyEmailPwdService
func NewVerifyEmailPwdService(ctx context.Context) *VerifyEmailPwdService {
	return &VerifyEmailPwdService{ctx: ctx}
}

// Run create note info
func (s *VerifyEmailPwdService) Run(req *micro_user.RpcVerifyEmailPwdReq) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
