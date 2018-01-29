package utils

import (
	"github.com/Hexilee/rady"
	"github.com/jinzhu/gorm"
)

func SetGormIfAutoRollback(db *gorm.DB) *gorm.DB {
	if rady.IsAutoRollback() {
		db = db.Begin()
	}
	return db
}
