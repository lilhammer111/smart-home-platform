package responder

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/bizerr"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"net/http"
	"reflect"
)

const (
	SuccessMessage = "message"
)

type FormattedResp struct {
	Success bool   `json:"success" example:"true"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Total   int    `json:"total,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func setCodeAndMsg(resp *FormattedResp, statusCode int) {
	resp.Code = statusCode
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	resp := FormattedResp{}
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	hlog.Debugf("*********************************** err is: %v", err)
	hlog.Debugf("*********************************** bizErr: %v", bizErr)
	// Though error is not biz err, it's also can be
	if !isBizErr {
		hlog.Debugf("*********************************** biz err type : %T", reflect.TypeOf(bizErr))
		hlog.Debugf("*********************************** origin err type : %T", reflect.TypeOf(err))
		resp.Message = http.StatusText(http.StatusInternalServerError)
		setCodeAndMsg(&resp, http.StatusBadGateway)
	} else {
		resp.Message = bizErr.BizMessage()
		switch bizErr.BizStatusCode() {
		case bizerr.CodeNotFound:
			resp.Code = http.StatusNotFound
		case bizerr.CodeAlreadyExists:
			resp.Code = http.StatusConflict
		case bizerr.CodeBadRequest:
			resp.Code = http.StatusBadRequest
		case bizerr.CodeExternalError:
			resp.Code = http.StatusBadGateway
		case bizerr.CodeAuthenticationFailed:
			resp.Code = http.StatusUnauthorized
		default:
			// do nothing
		}
	}

	c.JSON(code, resp)
}

func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	resp := FormattedResp{
		Success: true,
		Data:    data,
		Message: c.GetString(SuccessMessage),
	}
	switch string(c.Method()) {
	case http.MethodPost, http.MethodPut:
		resp.Code = http.StatusCreated
	case http.MethodDelete:
		resp.Code = http.StatusNoContent
	default:
		resp.Code = http.StatusOK
	}
	c.JSON(code, resp)
}
