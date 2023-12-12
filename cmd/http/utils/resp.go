package utils

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"go/types"
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
)

type respData struct {
	Success string
	Code    int16
	Message string
	Data    any
}

func formatData(src interface{}) *respData {
	// CopyFields copies fields from src to dst struct by matching field names.
	dst := &respData{}
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	fieldSuccess := srcVal.FieldByName("Meta").FieldByName("Success")
	fieldCode := srcVal.FieldByName("Meta").FieldByName("Code")
	fieldMessage := srcVal.FieldByName("Meta").FieldByName("Message")
	srcDataField := srcVal.FieldByName("Data")

	dstVal.FieldByName("Success").Set(fieldSuccess)
	dstVal.FieldByName("Code").Set(fieldCode)
	dstVal.FieldByName("Message").Set(fieldMessage)
	dstVal.FieldByName("Data").Set(srcDataField)

	return dst
}

type AlterResp interface {
	GetMeta() (v *standard.Resp)
}

type AlertListResp interface {
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
	//rd := formatData(data)

	type Response struct {
		Success bool
		Code    int16
		Message string
		Data    any
	}
	var response Response
	switch resp := data.(type) {
	case AlterResp:
		response.Success = resp.GetMeta().GetSuccess()
		response.Code = resp.GetMeta().GetCode()
		response.Message = resp.GetMeta().GetMessage()
		switch newResp := resp.(type) {
		case AlertInfoResp:
			response.Data = newResp.GetData()
		case AlertListResp:
			response.Data = newResp.GetData()
		}
	case types.Nil:
		// handle...
	default:
		// handle...
	}

	c.JSON(code, response)
}
