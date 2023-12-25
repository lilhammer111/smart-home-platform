package model

import (
	"git.zqbjj.top/lilhammer111/micro-kit/model"
)

type Alert struct {
	model.BaseModel
	DeviceId   int32  `gorm:"type:int unsigned;index;not null" json:"device_id"`
	IsOngoing  bool   `gorm:"type:bool;not null" json:"is_ongoing"`
	FirstAlarm string `gorm:"type:datetime;not null;index;comment:初次报警时间" json:"first_alarm"`
	LastAlarm  string `gorm:"type:datetime;not null;index;comment:最后报警时间" json:"last_alarm"`
	Count      int8   `gorm:"type:tinyint unsigned;not null;comment:报警次数" json:"count"`
	Level      int8   `gorm:"type:tinyint unsigned;not null" json:"level"`
	Desc       string `gorm:"type:varchar(255);not null" json:"desc"`
}
