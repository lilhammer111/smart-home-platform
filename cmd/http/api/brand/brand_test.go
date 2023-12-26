package brand

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetBrandList(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/brands/list", GetBrandList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/brands/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetBrandByCategory(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/categories/brands/list", GetBrandByCategory)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/categories/brands/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestAddNewBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/brands/add", AddNewBrand)
	w := ut.PerformRequest(h.Engine, "POST", "/api/products/brands/add", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/brands/update", UpdateBrand)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/brands/update", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDelteBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/brands/delete", DelteBrand)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/brands/delete", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
