package model

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetAllModels(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/models/list", GetAllModels)
	w := ut.PerformRequest(h.Engine, "GET", "/api/products/models/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestAddNewModel(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/models/add", AddNewModel)
	w := ut.PerformRequest(h.Engine, "POST", "/api/products/models/add", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteModel(t *testing.T) {
	h := server.Default()
	h.GET("/api/products/models/delete", DeleteModel)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/products/models/delete", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
