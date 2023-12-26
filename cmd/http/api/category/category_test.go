package category

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetCategoryList(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/categories/list", GetCategoryList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/categories/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetCategoryDetail(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/categories/detail", GetCategoryDetail)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/categories/detail", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestAddNewCategory(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/categories/add", AddNewCategory)
	w := ut.PerformRequest(h.Engine, "POST", "/api/products/categories/add", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateCategory(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/categories/update", UpdateCategory)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/categories/update", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteCategory(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/categories/delete", DeleteCategory)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/categories/delete", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
