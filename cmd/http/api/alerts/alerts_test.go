package alerts

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetAlertList(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/alerts", GetAlertList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/devices/alerts", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetAlertDetail(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/alerts/detail", GetAlertDetail)
	w := ut.PerformRequest(h.Engine, "GET", "/api/devices/alerts/detail", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateAlertInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/alerts", UpdateAlertInfo)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/devices/alerts", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUploadAlertInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/alerts", UploadAlertInfo)
	w := ut.PerformRequest(h.Engine, "POST", "/api/devices/alerts", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteAlert(t *testing.T) {
	h := server.Default()
	h.GET("/api/devices/alerts", DeleteAlert)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/devices/alerts", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
