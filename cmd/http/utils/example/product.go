package example

type ProductState struct {
	OnSale         bool `json:"on_sale" example:"true"`
	IsFreeShipping bool `json:"is_free_shipping" example:"false"`
	IsNew          bool `json:"is_new" example:"true"`
	IsHot          bool `json:"is_hot" example:"false"`
	IsRecommended  bool `json:"is_recommended" example:"true"`
}

type UpdateProductData struct {
	Id         int32    `json:"id" example:"1"`
	CategoryId int32    `json:"category_id" example:"10"`
	BrandId    int32    `json:"brand_id" example:"20"`
	ModelList  []string `json:"model_list" example:"model1,model2,model3"`
	Name       string   `json:"name" example:"Product Name"`
	Brief      string   `json:"brief" example:"Product Description"`
	Picture    string   `json:"picture" example:"https://example.com/picture.jpg"`
	Price      string   `json:"price" example:"99.99"`
	Showcase   []string `json:"showcase" example:"https://example.com/showcase1.jpg,https://example.com/showcase2.jpg"`
}

type AddProductData struct {
	CategoryId int32    `json:"category_id" example:"10"`
	BrandId    int32    `json:"brand_id" example:"20"`
	ModelList  []string `json:"model_list" example:"model1,model2,model3"`
	Name       string   `json:"name" example:"Product Name"`
	Brief      string   `json:"brief" example:"Product Description"`
	Picture    string   `json:"picture" example:"https://example.com/picture.jpg"`
	Price      float32  `json:"price" example:"99.99"`
	Showcase   []string `json:"showcase" example:"https://example.com/showcase1.jpg,https://example.com/showcase2.jpg"`
}
