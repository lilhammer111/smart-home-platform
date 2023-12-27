package example

type ProductState int8

const (
	OnSale ProductState = 1 << iota
	IsFreeShipping
	IsNew
	IsHot
	IsRecommended
)

type ProductData struct {
	Id         int32
	CategoryId int32 `gorm:"type:int unsigned;not null"`
	BrandId    int32 `gorm:"type:int unsigned;not null"`
	ModelId    int8  `gorm:"type:tinyint unsigned;not null"`

	Name    string `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_brand_model_name"`
	Brief   string `gorm:"type:varchar(100);not null"`
	Picture string `gorm:"type:varchar(200);not null;comment: oss url"`

	State    ProductState `gorm:"type:tinyint unsigned;not null"`
	Price    float32      `gorm:"type:float unsigned;index;not null"`
	RatingId int32        `gorm:"type:int unsigned;index;not null"`
	Showcase []string     `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}

type NewProductBody struct {
	CategoryId int32 `gorm:"type:int unsigned;not null"`
	BrandId    int32 `gorm:"type:int unsigned;not null"`
	ModelId    int8  `gorm:"type:tinyint unsigned;not null"`

	Name    string `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_brand_model_name"`
	Brief   string `gorm:"type:varchar(100);not null"`
	Picture string `gorm:"type:varchar(200);not null;comment: oss url"`

	State    ProductState `gorm:"type:tinyint unsigned;not null"`
	Price    float32      `gorm:"type:float unsigned;index;not null"`
	RatingId int32        `gorm:"type:int unsigned;index;not null"`
	Showcase []string     `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}
