package service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type VerifySmsCodeService struct {
	ctx context.Context
} // NewVerifySmsCodeService new VerifySmsCodeService
func NewVerifySmsCodeService(ctx context.Context) *VerifySmsCodeService {
	return &VerifySmsCodeService{ctx: ctx}
}

// Run create note info
func (s *VerifySmsCodeService) Run(req *micro_user.RpcVerifyCodeReq) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
