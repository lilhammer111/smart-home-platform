package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32 `gorm:"type:int unsigned;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
