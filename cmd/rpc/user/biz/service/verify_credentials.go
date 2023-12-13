package service

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type VerifyCredentialsService struct {
	ctx context.Context
} // NewVerifyCredentialsService new VerifyCredentialsService
func NewVerifyCredentialsService(ctx context.Context) *VerifyCredentialsService {
	return &VerifyCredentialsService{ctx: ctx}
}

// Run create note info
func (s *VerifyCredentialsService) Run(req *micro_user.RpcCredentialReq) (resp *common_rpc.RpcEmpty, err error) {
	// Finish your business logic.

	return
}
