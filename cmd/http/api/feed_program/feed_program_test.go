package feed_program

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetSelfProgramList(t *testing.T) {
	h := server.Default()
	h.GET("/api/programs/list", GetSelfProgramList)
	w := ut.PerformRequest(h.Engine, "GET", "/api/programs/list", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetProgramDetail(t *testing.T) {
	h := server.Default()
	h.GET("/api/programs/detail", GetProgramDetail)
	w := ut.PerformRequest(h.Engine, "GET", "/api/programs/detail", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateProgramInfo(t *testing.T) {
	h := server.Default()
	h.GET("/api/programs/update", UpdateProgramInfo)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/programs/update", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestCreateProgram(t *testing.T) {
	h := server.Default()
	h.GET("/api/programs/create", CreateProgram)
	w := ut.PerformRequest(h.Engine, "POST", "/api/programs/create", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestFeedNow(t *testing.T) {
	h := server.Default()
	h.GET("/api/programs/feed_now", FeedNow)
	w := ut.PerformRequest(h.Engine, "POST", "/api/programs/feed_now", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteProgram(t *testing.T) {
	h := server.Default()
	h.GET("/api/programs/delete", DeleteProgram)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/programs/delete", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
