package rorm

import (
	"github.com/Hexilee/rady"
	"github.com/Hexilee/rorm/sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

type OrmTest struct {
	rady.Testing
	SQLite *sqlite3.GormSQLite
}

func (orm *OrmTest) TestSQLite(t *testing.T) {
	DB := orm.SQLite.Begin()
	DB.Create(&User{Name: "xixi", Age: 18})
	NewUser := new(User)
	DB.Where(&User{Name: "xixi"}).Find(&NewUser)
	assert.Equal(t, 18, NewUser.Age)
	DB.Rollback()
}

func TestOrm(t *testing.T) {
	rady.CreateApplication(new(OrmRoot)).RunTest(t, new(OrmTest))
}
