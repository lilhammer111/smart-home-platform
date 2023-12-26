package example

type RespOk struct {
	Success bool   `example:"true"`
	Code    int    `example:"200"`
	Message string `example:"OK."`
}

type RespCreated struct {
	Success bool   `example:"true"`
	Code    int    `example:"201"`
	Message string `example:"Creating successes."`
}

type RespNoContent struct {
	Success bool   `example:"true"`
	Code    int    `example:"204"`
	Message string `example:"Deleting successes."`
}

type RespNotFound struct {
	Success bool   `example:"false"`
	Code    int    `example:"404"`
	Message string `example:"Not found."`
}

type RespInternal struct {
	Success bool   `example:"false"`
	Code    int    `example:"500"`
	Message string `example:"Internal server error."`
}

type RespBadRequest struct {
	Success bool   `example:"false"`
	Code    int    `example:"400"`
	Message string `example:"Bad request."`
}

type RespConflict struct {
	Success bool   `example:"false"`
	Code    int    `example:"409"`
	Message string `example:"Already exists."`
}

type RespUnauthorized struct {
	Success bool   `example:"false"`
	Code    int    `example:"401"`
	Message string `example:"Authentication fails."`
}

type Empty struct{}
