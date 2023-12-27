package example

type CategoryBrandData struct {
	Id         int32 `json:"id" example:"123"`
	CategoryId int32 `json:"categoryId" example:"10"`
	BrandId    int32 `json:"brandId" example:"5"`
}

type NewCategoryBrandBody struct {
	BrandId    int32   `json:"brandId" example:"5"`
	CategoryId []int32 `json:"categoryId" example:"[10, 11, 12]"`
}
