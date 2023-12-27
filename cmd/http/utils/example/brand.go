package example

type BrandData struct {
	Id   int32  `example:"1"`
	Name string `example:"xiaomi"`
	Logo string `example:"https://upload.wikimedia.org/wikipedia/commons/2/29/Xiaomi_logo.svg"`
}

type NewBrandBody struct {
	Name string `example:"xiaomi"`
	Logo string `example:"https://upload.wikimedia.org/wikipedia/commons/2/29/Xiaomi_logo.svg"`
}
