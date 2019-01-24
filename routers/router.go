package routers

import (
	"BeeMail/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/send", &controllers.SendMessageController{})
	beego.Router("/", &controllers.IncomingMessageController{})
}
