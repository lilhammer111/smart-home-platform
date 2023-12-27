package example

type CategoryBrandData struct {
	Id         int32 `json:"id" example:"123"`
	CategoryId int32 `json:"category_id" example:"10"`
	BrandId    int32 `json:"brand_id" example:"5"`
}

type NewCategoryBrandBody struct {
	BrandId    int32   `json:"brand_id" example:"5"`
	CategoryId []int32 `json:"category_id" example:"10,11,12"`
}
