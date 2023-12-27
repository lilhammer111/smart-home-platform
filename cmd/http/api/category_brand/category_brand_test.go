package category_brand

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

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

func TestBatchReduceCategoryBrand(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/category_brand/batch_reduce", BatchReduceCategoryBrand)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/category_brand/batch_reduce", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
