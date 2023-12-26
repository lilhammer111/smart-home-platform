package model

import "git.zqbjj.top/lilhammer111/micro-kit/model"

type ProductState int8

const (
	OnSale ProductState = 1 << iota
	IsFreeShipping
	IsNew
	IsHot
	IsRecommended
)

type Product struct {
	model.BaseModel
	CategoryId int32 `gorm:"type:int unsigned;not null"`
	BrandId    int32 `gorm:"type:int unsigned;not null"`

	Name    string `gorm:"type:varchar(50);not null"`
	ModelId int8   `gorm:"type:json;not null"`
	Brief   string `gorm:"type:varchar(100);not null"`
	Picture string `gorm:"type:varchar(200);not null;comment: oss url"`

	State    ProductState `gorm:"type:tinyint unsigned;not null"`
	Price    float32      `gorm:"type:float unsigned;index;not null"`
	Rating   int8         `gorm:"type:tinyint unsigned;index;not null"`
	Showcase []string     `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}
