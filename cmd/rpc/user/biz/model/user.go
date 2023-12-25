package model

import (
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/model"
	"gorm.io/gorm"
)

type User struct {
	model.BaseModel
	Username *string `gorm:"type:varchar(30);index;unique" json:"username"`
	Password *string `gorm:"type:varchar(255)" json:"password,omitempty"`
	Openid   *string `gorm:"type:varchar(50);index;unique" json:"openid"` // current length is 28
	Age      *int8   `gorm:"type:tinyint unsigned" json:"age"`
	Gender   *int8   `gorm:"type:tinyint unsigned" json:"gender"`
	Email    *string `gorm:"type:varchar(100);index;unique" json:"email"` // varchar(254) RFC5321 STANDARD
	Mobile   *string `gorm:"type:varchar(11);index;unique" json:"mobile"`
	Profile  *string `gorm:"type:varchar(255)" json:"profile"`
	Avatar   *string `gorm:"type:varchar(200)" json:"avatar"`
}

// BeforeDelete is a delete hook
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		return errors.New("admin user not allowed to delete")
	}
	return
}
