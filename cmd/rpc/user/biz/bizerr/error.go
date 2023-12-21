package bizerr

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

const (
	CodeUnknown int32 = 11 + iota
	CodeNotFound
	CodeBadRequest
	CodeAlreadyExists
	CodeInternalError
	CodeExternalError
	CodeAuthenticationFailed
)

const (
	MsgUnknown              = "rpc unknown error"
	MsgNotFound             = "rpc resource not found"
	MsgInternalError        = "rpc internal error"
	MsgExternalError        = "rpc external services unavailable"
	MsgAuthenticationFailed = "rpc authentication failed"
)

type RpcErr struct {
	code  int32
	msg   string
	extra map[string]string
}

// ServerErrorHandler convert errors that can be serialized
func ServerErrorHandler(ctx context.Context, err error) error {
	// if you want to get other rpc info, you can get rpcinfo first, like `ri := rpcinfo.GetRPCInfo(ctx)`
	// for example, get remote address: `remoteAddr := rpcinfo.GetRPCInfo(ctx).From().Address()`

	if errors.Is(err, kerrors.ErrBiz) {
		err = errors.Unwrap(err)
	}
	//if errCode, ok := geterror(err); ok {
	//	// for Thrift、KitexProtobuf
	//	return remote.NewTransError(errCode, err)
	//}
	return err
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

func NewAuthenticationError(err error) kerrors.BizStatusErrorIface {
	return newRpcError(CodeAuthenticationFailed, MsgAuthenticationFailed, err)
}
