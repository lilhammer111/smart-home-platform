package auth

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestSendSms(t *testing.T) {
	h := server.Default()
	h.GET("/api/auth/send_sms", SendSms)
	w := ut.PerformRequest(h.Engine, "GET", "/api/auth/send_sms", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestMobileRegister(t *testing.T) {
	h := server.Default()
	h.GET("/api/auth/mobile_register", MobileRegister)
	w := ut.PerformRequest(h.Engine, "POST", "/api/auth/mobile_register", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestMobileLogin(t *testing.T) {
	h := server.Default()
	h.GET("/api/auth/mobile_login", MobileLogin)
	w := ut.PerformRequest(h.Engine, "POST", "/api/auth/mobile_login", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestMiniProgLogin(t *testing.T) {
	h := server.Default()
	h.GET("/api/auth/mini_prog_login", MiniProgLogin)
	w := ut.PerformRequest(h.Engine, "POST", "/api/auth/mini_prog_login", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
