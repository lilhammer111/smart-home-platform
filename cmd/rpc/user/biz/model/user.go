package model

import (
	"git.zqbjj.top/pet/public-repo/model"
	"time"
)

type User struct {
	model.BaseModel
	Username string    `gorm:"type:varchar(30);index" json:"username"`
	Password string    `gorm:"type:varchar(255)" json:"password"`
	Openid   string    `gorm:"type:varchar(50)" json:"openid"` // current length is 28
	Age      int8      `gorm:"type:tinyint unsigned" json:"age"`
	Gender   int8      `gorm:"type:tinyint unsigned" json:"gender"`
	Email    string    `gorm:"type:varchar(100);index" json:"email"` // varchar(254) RFC5321 STANDARD
	Mobile   string    `gorm:"type:varchar(11);index" json:"mobile"`
	Profile  string    `gorm:"type:varchar(255)" json:"profile"`
	Avatar   string    `gorm:"type:varchar(200)" json:"avatar"`
	IsFrozen bool      `gorm:"type:bool;comment:whether user account frozen, 1 means user is frozen,0 for normal" json:"is_frozen"`
	ThawedAt time.Time `gorm:"type:datetime;comment:time when user accounts unfrozen" json:"thawed_at"`
}
