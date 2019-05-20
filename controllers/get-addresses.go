package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetAddressesController struct {
	beego.Controller
}

func (c *GetAddressesController) Get() {
	if !helpers.CheckIfLocalAddress(c.Ctx.Request.RemoteAddr) {
		c.Data["json"] = helpers.CreateResponse("Unauthorized")
		c.ServeJSON()
		return
	}
	db := *(database.GetInstance())
	var addressesMap []orm.Params
	_, err := db.QueryTable("mail").Distinct().Values(&addressesMap, "remote_address")
	if err != nil && err != orm.ErrNoRows {
		c.Data["json"] = helpers.CreateResponse("Failed to get messages")
		c.ServeJSON()
		return
	}
	var addresses []string
	for _, paramMap := range addressesMap {
		for _, address := range paramMap {
			addresses = append(addresses, address.(string))
		}
	}
	c.Data["json"] = addresses
	c.ServeJSON()
}
