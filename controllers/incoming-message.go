package controllers

import (
	"BeeMail/database"
	"BeeMail/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type IncomingMessageController struct {
	beego.Controller
}

func (c *IncomingMessageController) Post() {
	var msg models.Message
	json.Unmarshal(c.Ctx.Input.RequestBody, &msg)
	if msg.IsEmpty() {
		c.Data["json"] = map[string]string{"Response": "Message provided in improper format"}
	} else {
		c.Data["json"] = &msg
		db := *(database.GetInstance())
		db.Insert(&msg)
	}
	beego.Info("Received message:\n" + string(c.Ctx.Input.RequestBody))
	c.ServeJSON()
}
