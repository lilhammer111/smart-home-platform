package users

import (
	"context"
	"testing"

	resp "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/resp"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/users"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestDeregisterUserService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewDeregisterUserService(ctx, c)
	// init req and assert value
	req := &users.UserID{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
