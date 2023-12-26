package model

type Brand struct {
	Id   int32  `gorm:"type:int unsigned;primaryKey"`
	Name string `gorm:"type:varchar(20);unique;not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment: brand logo picture url"`
}

func (Brand) TableName() string {
	return "brands"
}
