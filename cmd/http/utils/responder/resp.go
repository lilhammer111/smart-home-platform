package responder

import (
	"context"
	rpcCode "git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"net/http"
)

const (
	SuccessMessage = "message"
	ErrorMessage   = "err_msg"
	ErrorCode      = "err_code"
)

type FormattedResp struct {
	Success bool   `json:"success" example:"true"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Total   int    `json:"total,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// SendErrResponse  pack error response.
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	hlog.Error(err)
	resp := FormattedResp{
		Success: false,
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
	// Determine if the error is a rpc call error or a business logic error
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if !isBizErr {
		if c.GetString(ErrorMessage) != "" {
			resp.Message = c.GetString(ErrorMessage)
		}

		if c.GetInt(ErrorCode) != 0 {
			resp.Code = c.GetInt(ErrorCode)
		}
	} else {
		resp.Message = bizErr.BizMessage()
		switch bizErr.BizStatusCode() {
		case rpcCode.BadRequest:
			resp.Code = http.StatusBadRequest // 400
		case rpcCode.AuthenticationFailed:
			resp.Code = http.StatusUnauthorized // 401
		case rpcCode.NotFound:
			resp.Code = http.StatusNotFound // 404
		case rpcCode.AlreadyExists:
			resp.Code = http.StatusConflict // 409
		case rpcCode.ExternalError, rpcCode.InternalError:
			resp.Code = http.StatusBadGateway // 502
		default:
			// do nothing
		}
	}
	c.JSON(code, resp)
}

// SendSuccessResponse  pack success response.
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	resp := FormattedResp{
		Success: true,
		Message: c.GetString(SuccessMessage),
		Data:    data,
	}
	switch string(c.Method()) {
	case http.MethodPost, http.MethodPut:
		resp.Code = http.StatusCreated // 201
	case http.MethodDelete:
		resp.Code = http.StatusNoContent // 204
	default:
		resp.Code = http.StatusOK // 200
	}
	c.JSON(code, resp)
}
