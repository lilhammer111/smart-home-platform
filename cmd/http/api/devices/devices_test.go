package devices

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetDeviceList(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices", GetDeviceList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/devices", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetDeviceDetail(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/detail", GetDeviceDetail)
	w := ut.PerformRequest(h.Engine, "GET", "/api/devices/detail", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateDeviceInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices", UpdateDeviceInfo)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/devices", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestBindDevice(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices", BindDevice)
	w := ut.PerformRequest(h.Engine, "POST", "/api/devices", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUnbindDevice(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices", UnbindDevice)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/devices", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
