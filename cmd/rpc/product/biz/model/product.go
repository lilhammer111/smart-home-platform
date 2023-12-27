package model

import (
	"git.zqbjj.top/lilhammer111/micro-kit/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

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
	ModelId    int8  `gorm:"type:tinyint unsigned;not null"`

	Name    string `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_brand_model_name"`
	Brief   string `gorm:"type:varchar(100);not null"`
	Picture string `gorm:"type:varchar(200);not null;comment: oss url"`

	State    ProductState `gorm:"type:tinyint unsigned;not null"`
	Price    float32      `gorm:"type:float unsigned;index;not null"`
	RatingId int32        `gorm:"type:int unsigned;index;not null"`
	Showcase []string     `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}

func (p Product) TableName() string {
	return "products"
}

func (p Product) AfterCreate(tx *gorm.DB) (err error) {
	rating := Rating{
		ProductId:     p.ID,
		TotalCustomer: 0,
		CurRating:     5,
	}
	err = tx.Create(&rating).Error
	if err != nil {
		klog.Errorf("failed to create rating record: %s", err)
		return err
	}

	p.RatingId = rating.Id
	return nil
}
