package main

import (
	"BeeMail/database"
	"BeeMail/helpers"
	_ "BeeMail/routers"
	"github.com/astaxie/beego"
)

func setup() {
	beego.BConfig.AppName = "BeeMail"
	beego.BConfig.ServerName = "BeeMail"
	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSPort = 1944
	beego.BConfig.Listen.HTTPSCertFile = "cryptography/BeeMail.crt"
	beego.BConfig.Listen.HTTPSKeyFile = "cryptography/BeeMail.key"
	beego.BConfig.RunMode = "prod"
	beego.BConfig.CopyRequestBody = true
}

func init() {
	helpers.CreateCertificateIfNotExists()
	database.GetInstance()
}

func main() {
	setup()
	beego.Run()
}
