package example

type ModelData struct {
	Id   int32  `example:"1"`
	Name string `example:"豪华旗舰款"`
}

type AddModelBody struct {
	Name string `example:"豪华旗舰款"`
}
