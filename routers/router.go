package routers

import (
	"BeeMail/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IncomingMessageController{})
}
