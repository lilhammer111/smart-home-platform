package handler

import (
	"context"
	"testing"

	model "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestAddNewModelService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewAddNewModelService(ctx, c)
	// init req and assert value
	req := &model.NewModel{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
