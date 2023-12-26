package model

type CategoryBrand struct {
	Id         int32 `gorm:"type:int unsigned;primaryKey"`
	CategoryId int32 `gorm:"type:int unsigned;not null;unique;index:idx_category_brand"`
	BrandId    int32 `gorm:"type:int unsigned;not null;unique;index:idx_category_brand"`
}
