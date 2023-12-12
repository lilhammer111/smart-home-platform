package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
)

type GetUsersByFilterService struct {
	ctx context.Context
} // NewGetUsersByFilterService new GetUsersByFilterService
func NewGetUsersByFilterService(ctx context.Context) *GetUsersByFilterService {
	return &GetUsersByFilterService{ctx: ctx}
}

// Run create note info
func (s *GetUsersByFilterService) Run(req *user.UsersFilterReq) (resp []*user.UserData, err error) {
	// Finish your business logic.

	return
}
