package example

type RatingData struct {
	ProductId     int32   `json:"product_id" example:"1"`
	TotalCustomer int32   `json:"total_customer" example:"250"`
	Rating        float32 `json:"rating" example:"4.5"`
}

type UpdateRatingData struct {
	ProductId int32   `json:"product_id" example:"1"`
	Rating    float64 `json:"rating" example:"4.5"`
}
