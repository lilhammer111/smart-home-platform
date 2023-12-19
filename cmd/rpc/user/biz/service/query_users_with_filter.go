package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type QueryUsersWithFilterService struct {
	ctx context.Context
}

// NewQueryUsersWithFilterService new QueryUsersWithFilterService
func NewQueryUsersWithFilterService(ctx context.Context) *QueryUsersWithFilterService {
	return &QueryUsersWithFilterService{ctx: ctx}
}

// Run create note info
func (s *QueryUsersWithFilterService) Run(req *user.UsersFilter) (resp []*user.UserInfo, err error) {
	// Finish your business logic.
	userInfos := make([]*model.User, 0)
	res := db.GetMysql().Scopes(scope.Paginate(req.Page, req.Limit)).Find(&userInfos)
	if res.Error != nil {
		klog.Errorf("failed to find all eligible user infos: %s", res.Error)
		return nil, bizerr.NewInternalError(res.Error)
	}
	if res.RowsAffected == 0 {
		// Proactively generating a rpc business error
		return nil, bizerr.NewNotFoundError(gorm.ErrRecordNotFound)
	}

	err = copier.CopyWithOption(&resp, userInfos, copier.Option{DeepCopy: true})
	if err != nil {
		klog.Errorf("failed to deep copy user infos to resp: %s", err)
		return nil, err
	}

	return
}
