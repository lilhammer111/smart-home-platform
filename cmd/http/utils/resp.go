package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"net/http"
)

const (
	SuccessMessage = "message"
	FailureMessage = "failureMessage"
	FailureCode    = "code"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	resp := &FormattedResp{}
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if !isBizErr {
		resp.Code = http.StatusInternalServerError
		resp.Message = http.StatusText(http.StatusInternalServerError)
		c.JSON(code, resp)
	} else {
		switch bizErr.BizStatusCode() {

		}
	}

	c.String(code, err.Error())
}

type FormattedResp struct {
	Success bool
	Code    int32
	Message string
	Data    any
}

func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	resp := &FormattedResp{
		Success: true,
		Message: c.GetString(SuccessMessage),
		Data:    data,
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
