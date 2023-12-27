package example

type CategoryData struct {
	Id       int32    `json:"id" example:"1"`
	Name     string   `json:"name" example:"smart pet feeder"`
	Brief    string   `json:"brief" example:"An automated feeder for pets"`
	Picture  string   `json:"picture" example:"https://example.com/pet_feeder.jpg"`
	Showcase []string `json:"showcase" example:"https://example.com/showcase1.jpg, https://example.com/showcase2.jpg"`
}

type NewCategoryBody struct {
	Name     string   `json:"name" example:"smart pet feeder"`
	Brief    string   `json:"brief" example:"An automated feeder for pets"`
	Picture  string   `json:"picture" example:"https://example.com/pet_feeder.jpg"`
	Showcase []string `json:"showcase" example:"https://example.com/showcase1.jpg, https://example.com/showcase2.jpg"`
}
