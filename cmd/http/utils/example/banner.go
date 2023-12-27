package example

type BannerData struct {
	Id          int8   `json:"id" example:"1"`
	Picture     string `json:"picture" example:"https://img.freepik.com/free-psd/bold-gradients-banner-template_23-2148819008.jpg?t=st=1703601665~exp=1703602265~hmac=84f6d13b72f5e743e64c58166cfe6323de3809450b6dfa9340cb80ab5e455a51"`
	ProductLink string `json:"product_link" example:"https://pet.service.com/prodcut-link"`
	Index       int8   `json:"index" example:"1"`
}

type NewBannerBody struct {
	Picture     string `json:"picture" example:"https://img.freepik.com/free-psd/bold-gradients-banner-template_23-2148819008.jpg?t=st=1703601665~exp=1703602265~hmac=84f6d13b72f5e743e64c58166cfe6323de3809450b6dfa9340cb80ab5e455a51"`
	ProductLink string `json:"product_link" example:"https://pet.service.com/prodcut-link"`
	Index       int8   `json:"index" example:"1"`
}
