package devices_status

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetDeviceStatus(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/status", GetDeviceStatus)
	w := ut.PerformRequest(h.Engine, "GET", "/api/devices/status", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestInitDeviceStatus(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/status", InitDeviceStatus)
	w := ut.PerformRequest(h.Engine, "POST", "/api/devices/status", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateDeviceStatus(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/status", UpdateDeviceStatus)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/devices/status", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
