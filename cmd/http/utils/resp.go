package utils

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
	"go/types"
)

type FormattedResp struct {
	Success bool
	Code    int16
	Message string
	Data    any
}

type StandardResp interface {
	GetMeta() (v *standard.Resp)
}

type ListResp interface {
	GetMeta() (v *standard.Resp)
	GetData() (v []int32)
}

type AlertInfoResp interface {
	GetMeta() (v *standard.Resp)
	GetData() (v *alert.AlertInfo)
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	var fr FormattedResp
	FormatResp(data, &fr)
	c.JSON(code, fr)
}

func FormatResp(data any, fr *FormattedResp) {
	switch resp := data.(type) {
	case StandardResp:
		fr.Success = resp.GetMeta().GetSuccess()
		fr.Code = resp.GetMeta().GetCode()
		fr.Message = resp.GetMeta().GetMessage()
		switch newResp := resp.(type) {
		case AlertInfoResp:
			fr.Data = newResp.GetData()
		case ListResp:
			fr.Data = newResp.GetData()
		}
	case types.Nil:
		// handle...
	default:
		// handle...
	}
}
