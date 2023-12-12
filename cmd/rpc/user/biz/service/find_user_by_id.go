package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
)

type FindUserByIDService struct {
	ctx context.Context
} // NewFindUserByIDService new FindUserByIDService
func NewFindUserByIDService(ctx context.Context) *FindUserByIDService {
	return &FindUserByIDService{ctx: ctx}
}

// Run create note info
func (s *FindUserByIDService) Run(req int32) (resp *user.UserData, err error) {
	// Finish your business logic.

	return
}
