package sqlite3

import (
	"github.com/Hexilee/rady"
	"github.com/Hexilee/rorm/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"reflect"
)

type (
	GormSQLiteConfig struct {
		rady.Configuration
		App *rady.Application
	}

	GormSQLiteParameter struct {
		rady.Parameter
		Path *string `value:"rady.sqlite3.path" default:"./rady.db"`
	}

	GormSQLite struct {
		rady.Database
		*gorm.DB
	}
)

func (g *GormSQLiteConfig) GetAutoMigrateSQLiteDB(params *GormSQLiteParameter) *GormSQLite {
	db, err := gorm.Open("sqlite3", *params.Path)
	if err != nil {
		g.App.Logger.Critical("Cannot access to sqlite \nError:\n%s", err.Error())
		os.Exit(1)
	}
	for _, entityType := range g.App.Entities {
		if entityType.Kind() == reflect.Ptr && entityType.Elem().Kind() == reflect.Struct {
			g.App.Logger.Debug("AutoMigrate: %s", entityType.String())
			db.AutoMigrate(reflect.New(entityType.Elem()).Interface())
		}
	}
	return &GormSQLite{DB: utils.SetGormIfAutoRollback(db)}
}
