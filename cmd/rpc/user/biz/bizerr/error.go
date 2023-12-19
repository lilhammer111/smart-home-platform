package bizerr

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

const (
	CodeUnknown int32 = iota
	CodeNotFound
	CodeBadRequest
	CodeAlreadyExists
	CodeInternalError
	CodeExternalError
)

const (
	MsgUnknown       = "rpc unknown error"
	MsgNotFound      = "rpc resource not found"
	MsgInternalError = "rpc internal error"
	MsgExternalError = "rpc external services unavailable"
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

func NewInternalError(err error) kerrors.BizStatusErrorIface {
	return newRpcError(CodeInternalError, MsgInternalError, err)

}

func NewNotFoundError(err error) kerrors.BizStatusErrorIface {
	return newRpcError(CodeNotFound, MsgNotFound, err)
}

func NewExternalError(err error) kerrors.BizStatusErrorIface {
	return newRpcError(CodeExternalError, MsgExternalError, err)
}
