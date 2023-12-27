package example

type RatingData struct {
	Id            int32   `json:"id" example:"1"`
	ProductId     int32   `json:"product_id" example:"1"`
	TotalCustomer int32   `json:"total_customer" example:"250"`
	CurRating     float32 `json:"cur_rating" example:"4.5"`
}

type RatingReqBody struct {
	ProductId int32   `json:"product_id" example:"1"`
	Rating    float64 `json:"rating" example:"4.5"`
}
