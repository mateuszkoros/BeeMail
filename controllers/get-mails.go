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
	var address string
	c.Ctx.Input.Bind(&address, "address")
	db := *(database.GetInstance())
	var mails []*models.Mail
	_, err := db.QueryTable("mail").Filter("remote_address", address).All(&mails)
	if err != nil && err != orm.ErrNoRows {
		c.Data["json"] = helpers.CreateResponse("Failed to get messages")
		c.ServeJSON()
		return
	}
	c.Data["json"] = mails
	c.ServeJSON()
}
