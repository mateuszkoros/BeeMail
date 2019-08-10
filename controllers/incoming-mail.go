package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"BeeMail/models"
	"github.com/astaxie/beego"
)

type IncomingMailController struct {
	beego.Controller
}

// Incoming mail endpoint accepts new messages and saves them in database.
func (c *IncomingMailController) Post() {
	mail := helpers.CreateMailFromHttpRequest(c.Ctx.Request)
	if mail.IsEmpty() {
		c.Data["json"] = helpers.CreateResponse("Mail provided in improper format")
		c.ServeJSON()
		return
	}
	mail.SetRemoteAddress(c.Ctx.Request.RemoteAddr)
	if mail.RemoteAddress == "" {
		c.Data["json"] = helpers.CreateResponse("Mail provided in improper format")
		c.ServeJSON()
		return
	}
	mail.Type = models.Incoming
	c.Data["json"] = helpers.CreateResponse("OK")
	db := *(database.GetInstance())
	_, err := db.Insert(&mail)
	helpers.CheckError(err)
	beego.Info("Received mail:\n" + string(c.Ctx.Input.RequestBody))
	c.ServeJSON()
}
