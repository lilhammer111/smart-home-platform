package model

import (
	"git.zqbjj.top/lilhammer111/micro-kit/model"
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
	OwnerId   int32 `gorm:"type:int unsigned;index;not null" json:"owner"`
	ProductId int32 `gorm:"type:smallint unsigned;not null" json:"type"`

	SerialNo *string `gorm:"type:varchar(255)" json:"serial_no"`
	Name     *string `gorm:"type:varchar(50);index" json:"name"`
	//Manufacturer    string `gorm:"type:varchar(50)" json:"manufacturer"`
	Desc            *string     `gorm:"type:varchar(255);comment:description" json:"desc"`
	State           DeviceState `gorm:"type:tinyint unsigned;not null" json:"status"`
	LocationId      *int16      `gorm:"type:smallint unsigned" json:"location_id"`
	HardwareVersion *string     `gorm:"type:varchar(50)" json:"hardware_version"`
	SoftwareVersion *string     `gorm:"type:varchar(50)" json:"software_version"`
}

type Location struct {
	Id    int16  `gorm:"type:smallint unsigned;primaryKey" json:"id"`
	Title string `gorm:"type:varchar(50);not null" json:"title"`
}
