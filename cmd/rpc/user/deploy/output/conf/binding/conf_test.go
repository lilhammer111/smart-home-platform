package binding

import (
	"testing"
)

const (
	username = "nacos"
)

func TestGetLocalConf(t *testing.T) {
	if GetLocalConf().Nacos.Username != username {
		t.Error("failed to get local conf")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestGetLocalConf", TestGetLocalConf)
}
