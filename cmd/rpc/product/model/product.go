package model

type ProductState int8

const (
	OnSale ProductState = 1 << iota
	IsFreeShipping
	IsNew
	IsHot
	IsRecommended
)

type Product struct {
	BaseModel
	CategoryId int32 `gorm:"type:int unsigned;not null"`
	BrandId    int32 `gorm:"type:int unsigned;not null"`

	Name    string   `gorm:"type:varchar(50);not null"`
	Model   []string `gorm:"type:json;not null"`
	Brief   string   `gorm:"type:varchar(100);not null"`
	Picture string   `gorm:"type:varchar(200);not null;comment: oss url"`

	State    ProductState `gorm:"type:tinyint unsigned;not null"`
	Price    float32      `gorm:"type:float;not null"`
	Rating   int8         `gorm:"type:tinyint unsigned;not null"`
	Showcase []string     `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}

type Category struct {
	BaseModel
	Name     string   `gorm:"type:varchar(20);not null"`
	Brief    string   `gorm:"type:varchar(100);not null"`
	Picture  string   `gorm:"type:varchar(200);not null;comment: oss url"`
	Showcase []string `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment: brand logo picture url"`
}

type Banner struct {
	BaseModel
	Picture string `gorm:"type:varchar(200);not null;comment: oss url"`
	Url     string `gorm:"type:varchar(200);not null"`
	Index   int8   `gorm:"type:tinyint unsigned;unique;not null"`
}
