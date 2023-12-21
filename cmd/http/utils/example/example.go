package example

type RespOk struct {
	Success bool   `example:"true"`
	Code    int    `example:"200"`
	Message string `example:"ok"`
	Total   int    `example:"1"`
	Data    UserData
}

type RespCreated struct {
	Success bool   `example:"true"`
	Code    int    `example:"201"`
	Message string `example:"registry ok"`
	Total   int    `example:"1"`
	Data    UserData
}

type UserData struct {
	Id       *int32  `example:"1"`
	Age      *int8   `example:"18"`
	Gender   *int8   `example:"1"`
	Mobile   *string `example:"19535876981"`
	Profile  *string `example:"hello,world"`
	Username *string `example:"lilhammer111"`
	Password *string `example:"12345678"`
	Email    *string `example:"wwwwwdemon@gmail.com"`
	Avatar   *string `example:""`
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
