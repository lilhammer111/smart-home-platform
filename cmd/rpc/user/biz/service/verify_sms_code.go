package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
)

type VerifySmsCodeService struct {
	ctx context.Context
} // NewVerifySmsCodeService new VerifySmsCodeService
func NewVerifySmsCodeService(ctx context.Context) *VerifySmsCodeService {
	return &VerifySmsCodeService{ctx: ctx}
}

// Run create note info
func (s *VerifySmsCodeService) Run(mobile string, smsCode string) (resp bool, err error) {
	// Finish your business logic.
	res, err := db.GetRedis().Get(context.Background(), mobile).Result()
	if err != nil {
		return false, bizerr.NewMysqlErr(err)
	}

	return res == smsCode, nil
}
