package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CList []string

func (g CList) Value() (driver.Value, error) {
	b, err := json.Marshal(g)
	if err != nil {
		klog.Errorf("failed to marshal string slice type: %s", err)
		return nil, err
	}
	return b, nil
}

func (g *CList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
