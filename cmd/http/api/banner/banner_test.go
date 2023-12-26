package banner

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetAllBanners(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/banners/list", GetAllBanners)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/banners/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestAddNewBanner(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/banners/add", AddNewBanner)
	w := ut.PerformRequest(h.Engine, "POST", "/api/products/banners/add", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateBanner(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/banners/update", UpdateBanner)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/banners/update", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDelteBanner(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/banners/delete", DelteBanner)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/banners/delete", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
