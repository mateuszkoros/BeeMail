package main

import (
	"BeeMail/database"
	_ "BeeMail/routers"
	"github.com/astaxie/beego"
)

func init() {
	database.GetInstance()
}
func main() {
	beego.Run()
}
