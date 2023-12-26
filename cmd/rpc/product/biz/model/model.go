package model

type Model struct {
	Id   int32  `gorm:"type:tinyint unsigned;primaryKey"`
	Name string `gorm:"type:varchar(50);unique"`
}

func (Model) TableName() string {
	return "models"
}
