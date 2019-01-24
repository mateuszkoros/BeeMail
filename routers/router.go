package routers

import (
	"BeeMail/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/send", &controllers.SendMailController{})
	beego.Router("/", &controllers.IncomingMailController{})
}
