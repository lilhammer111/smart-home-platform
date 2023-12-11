// Code generated by hertz generator. DO NOT EDIT.

package devices_status

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	devices_status "git.zqbjj.top/pet/services/cmd/http/api/devices_status"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_devices := _api.Group("/devices", _devicesMw()...)
			_devices.GET("/status", append(_getdevicestatusMw(), devices_status.GetDeviceStatus)...)
			_devices.POST("/status", append(_initdevicestatusMw(), devices_status.InitDeviceStatus)...)
			_devices.PUT("/status", append(_updatedevicestatusMw(), devices_status.UpdateDeviceStatus)...)
		}
	}
}
