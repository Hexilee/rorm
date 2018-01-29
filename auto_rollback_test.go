package rorm

import (
	"github.com/Hexilee/rady"
	"github.com/Hexilee/rorm/sqlite3"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type AutoRollBackTest struct {
	rady.Testing
	SQLite *sqlite3.GormSQLite
}

func (orm *AutoRollBackTest) TestSQLite(t *testing.T) {
	if rady.IsAutoRollback() {
		DB := orm.SQLite
		DB.Create(&User{Name: "xixi", Age: 18})
		NewUser := new(User)
		DB.Where(&User{Name: "xixi"}).Find(&NewUser)
		assert.Equal(t, 18, NewUser.Age)
		DB.Rollback()
	} else {
		DB := orm.SQLite.Begin()
		NewUser := new(User)
		DB.Where(&User{Name: "xixi"}).Find(&NewUser)
		assert.Equal(t, 0, NewUser.Age)
		DB.Create(&User{Name: "xixi", Age: 18})
		DB.Where(&User{Name: "xixi"}).Find(&NewUser)
		assert.Equal(t, 18, NewUser.Age)
		DB.Rollback()
	}
}

func TestAutoRollBack(t *testing.T) {
	rady.CreateApplication(new(OrmRoot)).RunTest(t, new(AutoRollBackTest))
	os.Setenv(rady.AutoRollbackEnv, rady.AutoRollback)
	rady.CreateApplication(new(OrmRoot)).RunTest(t, new(AutoRollBackTest))
	rady.ResetEnv(rady.AutoRollbackEnv)
}
