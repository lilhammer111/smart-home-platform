package model

import "gorm.io/gorm"

type BaseModel struct {
	Id        int32
	CreatedAt int32          `gorm:"autoCreateTime:milli"`
	UpdatedAt int32          `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
