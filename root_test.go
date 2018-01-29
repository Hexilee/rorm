package rorm

import "github.com/Hexilee/rorm/sqlite3"

type OrmRoot struct {
	*sqlite3.GormSQLiteConfig
	*OrmStorage
}
