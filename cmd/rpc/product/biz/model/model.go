package model

type Model struct {
	Id        int32  `gorm:"type:int unsigned;primaryKey"`
	ProductId int32  `gorm:"type:int unsigned;uniqueIndex:idx_product_id_name"`
	Name      string `gorm:"type:varchar(50);uniqueIndex:idx_product_id_name"`
}

func (Model) TableName() string {
	return "models"
}
