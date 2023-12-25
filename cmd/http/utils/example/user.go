package example

type AuthData struct {
	Token     string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDM0ODc1NTEsIm9yaWdfaWF0IjoxNzAzNDgzOTUxLCJ1aWQiOjN9.J5sOSjOPIgifaTpGIqzfZV3vi1ZRD6WnOJV3fok_ltk" json:"token,omitempty"`
	ExpiredAt string `example:"2023-12-25T14:59:11.322480873+08:00" json:"expired_at,omitempty"`
}

type UserData struct {
	Id       *int32  `example:"1" json:"id,omitempty"`
	Age      *int8   `example:"18" json:"age,omitempty"`
	Gender   *int8   `example:"1" json:"gender,omitempty"`
	Mobile   *string `example:"19535876981" json:"mobile,omitempty"`
	Profile  *string `example:"hello,world" json:"profile,omitempty"`
	Username *string `example:"lilhammer111" json:"username,omitempty"`
	Password *string `example:"12345678" json:"password,omitempty"`
	Email    *string `example:"wwwwwdemon@gmail.com" json:"email,omitempty"`
	Avatar   *string `example:"" json:"avatar,omitempty"`
}

type PwdLoginBody struct {
	Username string `example:"demon_wang"`
	//    2: required string Email (api.body="email", api.vd="regexp('^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$')");
	Password string `example:"12345678"`
}

type MobileLoginBody struct {
	Mobile  string `example:"19535876981"`
	SmsCode string `example:"123456"`
}

type MiniProgLoginBody struct {
	WxCode string `example:"wx1234567890abcdef1234567890abcdef"`
}

type MobileRegisterBody struct {
	Mobile   string `example:"19535876981"`
	SmsCode  string `example:"123456"`
	Username string `example:"demon_wang"`
	Password string `example:"12345678"`
}

type UsernameRegisterBody struct {
	Username string `example:"demon_wang"`
	Password string `example:"12345678"`
}
