package category_brand

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetCategoryBrandList(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/category_brand/list", GetCategoryBrandList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/category_brand/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestBatchAddCategoryBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/category_brand/batch_add", BatchAddCategoryBrand)
	w := ut.PerformRequest(h.Engine, "POST", "/api/products/category_brand/batch_add", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateCategoryBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/category_brand/update", UpdateCategoryBrand)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/products/category_brand/update", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteCategoryByBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/category_brand/delete_category", DeleteCategoryByBrand)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/category_brand/delete_category", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
