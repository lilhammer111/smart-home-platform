// Code generated by hertz generator.

package device

import (
	"git.zqbjj.top/pet/services/cmd/http/mw"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _devicesMw() []app.HandlerFunc {
	handlerFunc := make([]app.HandlerFunc, 0)
	handlerFunc = append(handlerFunc, mw.JwtMiddleware.MiddlewareFunc())
	return handlerFunc
}

func _binddeviceMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getdevicedetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getdevicelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _unbinddeviceMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatedeviceinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}