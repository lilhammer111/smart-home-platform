package service

import (
	"context"
	"errors"
	"fmt"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateUserService) Run(req *user.UserInfo) (resp *user.UserInfo, err error) {
	// encrypt user password
	userInfo := model.User{}
	if err = copier.Copy(&userInfo, req); err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	klog.Debugf("************** user info : %+v", userInfo)

	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(*userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	pwd := string(encryptedPwd)
	userInfo.Password = &pwd

	if err = db.GetMysql().Create(&userInfo).Error; err != nil {
		klog.Error(err)
		var mysqlErr *mysql.MySQLError
		ok := errors.As(err, &mysqlErr)
		if ok && mysqlErr.Number == 1062 {
			return nil, kerrors.NewBizStatusError(code.AlreadyExists, fmt.Sprintf("account of username %s already exists", req.Username))
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &user.UserInfo{}
	err = copier.Copy(resp, &userInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return
}
