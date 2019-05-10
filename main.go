package main

import (
	"BeeMail/database"
	"BeeMail/helpers"
	_ "BeeMail/routers"
	"github.com/astaxie/beego"
)

func init() {
	helpers.CreateCertificateIfNotExists()
	database.GetInstance()
}
func main() {
	beego.Run()
}
