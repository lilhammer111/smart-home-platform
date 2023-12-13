package user

import (
	"context"
	"testing"

	user "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetUserListService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetUserListService(ctx, c)
	// init req and assert value
	req := &user.UsersFilter{}
	resp, err := s.Do(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
