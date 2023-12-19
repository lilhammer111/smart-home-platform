package bizerr

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

const (
	CodeUnknown = iota + 1
	CodeNotFound
	CodeBadRequest
	CodeAlreadyExists
	CodeInternal
	CodeBadGateway
)

const (
	MsgUnknown  = "rpc unknown error"
	MsgNotFound = "rpc resource not found"
	MsgInternal = "rpc internal error"
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
	return fmt.Sprintf("rpc error: { code: %d, msg: { %s } }", e.code, e.msg)
}

func newRpcError(code int32, msg string, err error) *RpcErr {
	if err == nil {
		return &RpcErr{
			code: CodeUnknown,
			msg:  MsgUnknown,
		}
	}

	return &RpcErr{
		code: code,
		msg:  fmt.Sprintf("%s: %s", msg, err),
	}
}

func NewInternalErr(err error) kerrors.BizStatusErrorIface {
	return newRpcError(CodeInternal, MsgInternal, err)

}

func NewNotFoundError(err error) kerrors.BizStatusErrorIface {
	return newRpcError(CodeNotFound, MsgNotFound, err)
}
