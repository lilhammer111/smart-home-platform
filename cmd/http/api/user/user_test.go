package user

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetCurUserInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/users/current", GetCurUserInfo)
	w := ut.PerformRequest(h.Engine, "GET", "/api/users/current", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetUserList(t *testing.T) {
	h := server.Default()
	h.GET("/api/users/list", GetUserList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/users/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetUserDetail(t *testing.T) {
	h := server.Default()
	h.GET("/api/users/detail", GetUserDetail)
	w := ut.PerformRequest(h.Engine, "GET", "/api/users/detail", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateUserInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/users/update", UpdateUserInfo)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/users/update", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeregisterUser(t *testing.T) {
	h := server.Default()
	h.GET("/api/users/deregister", DeregisterUser)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/users/deregister", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
