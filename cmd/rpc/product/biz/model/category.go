package model

import "git.zqbjj.top/pet/services/cmd/rpc/product/biz/model/ctype"

type Category struct {
	Id       int32       `gorm:"type:int unsigned;primaryKey"`
	Name     string      `gorm:"type:varchar(20);unique;not null"`
	Brief    string      `gorm:"type:varchar(100);not null"`
	Picture  string      `gorm:"type:varchar(200);not null;comment: oss url"`
	Showcase ctype.CList `gorm:"type:varchar(1000);not null;comment:provide a range of product illustrations also as product detail"`
}

func (Category) TableName() string {
	return "categories"
}
