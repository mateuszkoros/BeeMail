package database

import (
	"BeeMail/helpers"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var (
	instance orm.Ormer
	once     sync.Once
)

func GetInstance() *orm.Ormer {
	once.Do(func() {
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", "./messages.db")
		instance = orm.NewOrm()
		instance.Using("default")
		err := orm.RunSyncdb("default", false, true)
		helpers.CheckError(err)
	})
	return &instance
}
