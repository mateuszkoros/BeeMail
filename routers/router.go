package routers

import (
	"BeeMail/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/send", &controllers.SendMailController{})
	beego.Router("/get", &controllers.GetMailsController{})
	beego.Router("/addresses", &controllers.GetAddressesController{})
	beego.Router("/delete", &controllers.DeleteMailController{})
	beego.Router("/", &controllers.IncomingMailController{})
}
