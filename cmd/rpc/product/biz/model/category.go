package model

type Category struct {
	Id       int32    `gorm:"type:int unsigned;primaryKey"`
	Name     string   `gorm:"type:varchar(20);not null"`
	Brief    string   `gorm:"type:varchar(100);not null"`
	Picture  string   `gorm:"type:varchar(200);not null;comment: oss url"`
	Showcase []string `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}
