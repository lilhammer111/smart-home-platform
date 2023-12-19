package responder

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"net/http"
)

const (
	SuccessMessage = "message"
	ErrorMessage   = "error"
)

type FormattedResp struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Total   int    `json:"total,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func setCodeAndMsg(resp *FormattedResp, statusCode int) {
	resp.Code = statusCode
	resp.Message = http.StatusText(statusCode)
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	resp := FormattedResp{
		Message: c.GetString(ErrorMessage),
	}
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if !isBizErr {
		setCodeAndMsg(&resp, http.StatusBadGateway)
	}
	switch bizErr.BizStatusCode() {
	case bizerr.CodeNotFound:
		setCodeAndMsg(&resp, http.StatusNotFound)
	case bizerr.CodeAlreadyExists:
		setCodeAndMsg(&resp, http.StatusConflict)
	case bizerr.CodeBadRequest:
		setCodeAndMsg(&resp, http.StatusBadRequest)
	case bizerr.CodeExternalError:
		setCodeAndMsg(&resp, http.StatusBadGateway)
	case bizerr.CodeAuthenticationFailed:
		setCodeAndMsg(&resp, http.StatusUnauthorized)
	default:
		setCodeAndMsg(&resp, http.StatusInternalServerError)
	}
	c.JSON(code, resp)
}

func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	//resp := &FormattedResp{
	//	Success: true,
	//	Message: c.GetString(SuccessMessage),
	//	Data:    data,
	//}
	resp := FormattedResp{
		Success: true,
		Data:    data,
	}

	switch string(c.Method()) {
	case http.MethodPost, http.MethodPut:
		setCodeAndMsg(&resp, http.StatusCreated)
	case http.MethodDelete:
		setCodeAndMsg(&resp, http.StatusNoContent)
	default:
		setCodeAndMsg(&resp, http.StatusOK)
	}
	c.JSON(code, resp)
}
