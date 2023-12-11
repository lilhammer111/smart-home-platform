package service

import (
	"context"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/dto/kitex_gen/users"
)

type FindUserService struct {
	ctx context.Context
} // NewFindUserService new FindUserService
func NewFindUserService(ctx context.Context) *FindUserService {
	return &FindUserService{ctx: ctx}
}

// Run create note info
func (s *FindUserService) Run(req int32) (resp *users.UserData, err error) {
	// Finish your business logic.

	return
}
