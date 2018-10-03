package database

import (
	"BeeMail/helpers"
	"github.com/asdine/storm"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var (
	instance storm.DB
	once     sync.Once
)

func GetInstance() *storm.DB {
	once.Do(func() {
		// orm.RegisterDriver("sqlite3", orm.DRSqlite)
		// orm.RegisterDataBase("default", "sqlite3", "./messages.db")
		// instance = orm.NewOrm()
		// instance.Using("default")
		// err := orm.RunSyncdb("default", false, true)
		instance, err := storm.Open("messages.db")
		defer instance.Close()
		helpers.CheckError(err)
	})
	return &instance
}
