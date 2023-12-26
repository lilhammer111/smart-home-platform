package utils

import (
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/go-sql-driver/mysql"
)

func CheckResourceConflict(err error, respMsg string) (bool, kerrors.BizStatusErrorIface) {
	var mySQLErr *mysql.MySQLError
	if errors.As(err, &mySQLErr) && mySQLErr.Number == 1062 {
		return true, kerrors.NewBizStatusError(code.AlreadyExists, respMsg)
	}
	return false, nil
}
