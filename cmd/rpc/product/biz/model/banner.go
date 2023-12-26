package model

type Banner struct {
	Id          int8   `gorm:"type:int unsigned;primaryKey"`
	Picture     string `gorm:"type:varchar(200);not null;comment: oss url"`
	ProductLink string `gorm:"type:varchar(200);not null"`
	Index       int8   `gorm:"type:tinyint unsigned;unique;not null"`
}
