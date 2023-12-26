package product

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetProductList(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/list", GetProductList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetProductDetail(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/detail", GetProductDetail)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/detail", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestAddNewProduct(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/add", AddNewProduct)
	w := ut.PerformRequest(h.Engine, "POST", "/api/products/add", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateProductBasicInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/update_basis/", UpdateProductBasicInfo)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/update_basis/", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateProdShowcase(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/update_showcase", UpdateProdShowcase)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/update_showcase", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateRating(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/update_rating", UpdateRating)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/update_rating", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteProduct(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/delete", DeleteProduct)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/delete", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
