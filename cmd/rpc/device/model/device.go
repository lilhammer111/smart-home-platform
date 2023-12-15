package model

import (
	"git.zqbjj.top/pet/public-repo/pkg/model"
	"time"
)

// DeviceState is progressive, if connected, it must be normal and powered on
// Normal是否故障，PoweredOn是否开机，Connected是否连接网络，Working是否空闲
type DeviceState int8

const (
	Normal DeviceState = 1 << iota
	PoweredOn
	Connected
	Working
)

type Device struct {
	model.BaseModel
	OwnerId   int32 `gorm:"type:int unsigned;index" json:"owner"`
	ProductId int16 `gorm:"type:smallint unsigned" json:"type"`

	SerialNo string `gorm:"type:varchar(255)" json:"serial_no"`
	Nickname string `gorm:"type:varchar(50);index" json:"nickname"`
	//Manufacturer    string `gorm:"type:varchar(50)" json:"manufacturer"`
	Desc            string      `gorm:"type:varchar(255);comment:description" json:"desc"`
	State           DeviceState `gorm:"type:tinyint unsigned;not null" json:"status"`
	LocationId      int16       `gorm:"type:smallint unsigned" json:"location"`
	HardwareVersion string      `gorm:"type:varchar(50)" json:"hardware_version"`
	SoftwareVersion string      `gorm:"type:varchar(50)" json:"software_version"`
}

type LocationEnum struct {
	Id   int16  `gorm:"type:smallint unsigned;primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);not null" json:"name"`
}

type DeviceAlert struct {
	model.BaseModel
	DeviceId   int32     `gorm:"type:int unsigned;not null" json:"device_id"`
	FirstAlarm time.Time `gorm:"type:datetime;not null;comment:初次报警时间" json:"first_alarm"`
	LastAlarm  time.Time `gorm:"type:datetime;not null;comment:最后报警时间" json:"last_alarm"`
	Times      int8      `gorm:"type:tinyint unsigned;not null;comment:报警次数" json:"times"`
	Level      int8      `gorm:"type:tinyint unsigned;not null" json:"level"`
	Desc       string    `gorm:"type:varchar(255);not null" json:"desc"`
}
