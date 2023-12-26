package example

type RespOk struct {
	Success bool   `example:"true"`
	Code    int    `example:"200"`
	Message string `example:"ok"`
}

type RespCreated struct {
	Success bool   `example:"true"`
	Code    int    `example:"201"`
	Message string `example:"registry ok"`
}

type RespNotFound struct {
	Success bool   `example:"false"`
	Code    int    `example:"404"`
	Message string `example:"not found"`
}

type RespInternal struct {
	Success bool   `example:"false"`
	Code    int    `example:"500"`
	Message string `example:"internal error"`
}

type RespBadRequest struct {
	Success bool   `example:"false"`
	Code    int    `example:"400"`
	Message string `example:"bad request"`
}

type RespConflict struct {
	Success bool   `example:"false"`
	Code    int    `example:"409"`
	Message string `example:"account already exists"`
}

type RespUnauthorized struct {
	Success bool   `example:"false"`
	Code    int    `example:"401"`
	Message string `example:"authentication fails"`
}

type Empty struct{}
