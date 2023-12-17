package bizerr

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

const (
	unknownErrCode       = 1000
	businessLogicErrCode = 1001
	internalErrCode      = 1002
	thirdPartyErrCode    = 1003

	redisDBErrCode = 2001
	mysqlDBErrCode = 2001
)

const (
	unknownErr           = "unknown error"
	thirdPartyServiceErr = "third party service error"
	internalErr          = "internal error"

	redisDBErr = "operate redis error"
	mysqlDBErr = "operate mysql error"
)

type RpcErr struct {
	code  int32
	msg   string
	extra map[string]string
}

func (e *RpcErr) BizStatusCode() int32 {
	return e.code
}

func (e *RpcErr) BizMessage() string {
	return e.msg
}

func (e *RpcErr) BizExtra() map[string]string {
	return e.extra
}

func (e *RpcErr) Error() string {
	return fmt.Sprintf("rpc biz error: { code: %d, msg: { %s } }", e.code, e.msg)
}

func newRpcError(code int32, msg string, err error) *RpcErr {
	if err == nil {
		return &RpcErr{
			code: unknownErrCode,
			msg:  unknownErr,
		}
	}

	return &RpcErr{
		code: code,
		msg:  fmt.Sprintf("%s: %s", msg, err),
	}
}

func NewThirdPartyErr(err error) kerrors.BizStatusErrorIface {
	return newRpcError(thirdPartyErrCode, thirdPartyServiceErr, err)
}

func NewInternalErr(err error) kerrors.BizStatusErrorIface {
	return newRpcError(internalErrCode, internalErr, err)

}

func NewRedisErr(err error) kerrors.BizStatusErrorIface {
	return newRpcError(redisDBErrCode, redisDBErr, err)
}

func NewMysqlErr(err error) kerrors.BizStatusErrorIface {
	return newRpcError(mysqlDBErrCode, mysqlDBErr, err)
}

func NewBusinessLogicErr(err error, msg string) kerrors.BizStatusErrorIface {
	return newRpcError(businessLogicErrCode, msg, err)
}
