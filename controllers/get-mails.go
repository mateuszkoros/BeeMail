package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"BeeMail/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetMailsController struct {
	beego.Controller
}

func (c *GetMailsController) Get() {
	if !helpers.CheckIfLocalAddress(c.Ctx.Request.RemoteAddr) {
		c.Data["json"] = helpers.CreateResponse("Unauthorized")
		c.ServeJSON()
		return
	}
	var address string
	err := c.Ctx.Input.Bind(&address, "address")
	if err != nil {
		c.Data["json"] = helpers.CreateResponse("Please provide address")
		c.ServeJSON()
		return
	}
	db := *(database.GetInstance())
	var mails []*models.Mail
	_, err = db.QueryTable("mail").Filter("remote_address", address).All(&mails)
	if err != nil && err != orm.ErrNoRows {
		c.Data["json"] = helpers.CreateResponse("Failed to get messages")
		c.ServeJSON()
		return
	}
	c.Data["json"] = mails
	c.ServeJSON()
}
