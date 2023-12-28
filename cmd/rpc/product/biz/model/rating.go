package model

type ProductRating struct {
	ProductId     int32   `gorm:"type:int unsigned;primaryKey"`
	TotalRating   float32 `gorm:"type:float(9,1) unsigned;not null"`
	TotalCustomer int32   `gorm:"type:int unsigned;not null"`
}
