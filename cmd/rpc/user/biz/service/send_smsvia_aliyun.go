package service

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type SendSMSViaAliyunService struct {
	ctx context.Context
} // NewSendSMSViaAliyunService new SendSMSViaAliyunService
func NewSendSMSViaAliyunService(ctx context.Context) *SendSMSViaAliyunService {
	return &SendSMSViaAliyunService{ctx: ctx}
}

// Run create note info
func (s *SendSMSViaAliyunService) Run(req *micro_user.RpcSmsReq) (resp *common_rpc.RpcEmpty, err error) {
	// Finish your business logic.

	return
}
