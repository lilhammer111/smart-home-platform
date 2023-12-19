package scope

import (
	"gorm.io/gorm"
)

const (
	defaultPage  = 1
	defaultLimit = 10
)

func Paginate(page, limit *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var p, l int32
		if page == nil {
			p = defaultPage
		} else {
			if *page < 0 {
				p = defaultPage
			}
			p = *page
		}

		if limit == nil {
			l = defaultLimit
		} else {
			if *limit > 100 {
				l = 100
			} else if *limit < 0 {
				l = 1
			} else {
				l = *limit
			}
		}
		offset := (p - 1) * l
		return db.Offset(int(offset)).Limit(int(l))
	}
}
