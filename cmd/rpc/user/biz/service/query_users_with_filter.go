package service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"github.com/cloudwego/kitex/pkg/kerrors"

	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
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
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Debug("rows affected is 0")
		// Proactively generating a rpc business error
		return nil, kerrors.NewBizStatusError(code.NotFound, "no eligible users")
	}

	resp = make([]*user.UserInfo, 0)
	err = copier.CopyWithOption(&resp, userInfos, copier.Option{DeepCopy: true})
	if err != nil {
		klog.Errorf("failed to deep copy user infos to resp: %s", err)
		return nil, err
	}

	return
}
