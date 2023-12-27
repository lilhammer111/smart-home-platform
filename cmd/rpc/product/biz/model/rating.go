package model

type Rating struct {
	Id            int32   `gorm:"type:int unsigned;primaryKey"`
	ProductId     int32   `gorm:"type:int unsigned;not null;unique"`
	TotalCustomer int32   `gorm:"type:int unsigned;not null"`
	CurRating     float32 `gorm:"type:float unsigned;not null"`
}
