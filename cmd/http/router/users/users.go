// Code generated by hertz generator. DO NOT EDIT.

package users

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	users "git.zqbjj.top/pet/services/cmd/http/api/users"
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
		_api.PUT("/users", append(_updateuserinfoMw(), users.UpdateUserInfo)...)
		_api.DELETE("/users", append(_deregisteruserMw(), users.DeregisterUser)...)
		_api.GET("/users", append(_usersMw(), users.GetUserList)...)
		_users := _api.Group("/users", _usersMw()...)
		_users.GET("/detail", append(_getuserdetailMw(), users.GetUserDetail)...)
	}
}
