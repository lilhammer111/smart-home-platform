// Code generated by hertz generator.

package user

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

func _usersMw() []app.HandlerFunc {
	handlerFunc := make([]app.HandlerFunc, 0)
	handlerFunc = append(handlerFunc, mw.JwtMiddleware.MiddlewareFunc())
	return handlerFunc
}

func _deregisteruserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserdetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcuruserinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}
