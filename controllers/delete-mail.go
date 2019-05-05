package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"BeeMail/models"
	"github.com/astaxie/beego"
	"strconv"
)

type DeleteMailController struct {
	beego.Controller
}

func (c *DeleteMailController) Delete() {
	db := *(database.GetInstance())
	if len(c.Ctx.Request.Form["Id"]) == 0 {
		c.Data["json"] = helpers.CreateResponse("Please specify message to delete")
		c.ServeJSON()
		return
	}
	idAsString := c.Ctx.Request.Form["Id"][0]
	id, err := strconv.ParseUint(idAsString, 10, 0)
	if err != nil {
		c.Data["json"] = helpers.CreateResponse("Improper message id specified")
		c.ServeJSON()
		return
	}
	num, err := db.Delete(&models.Mail{Id: uint(id)})
	if err != nil {
		c.Data["json"] = helpers.CreateResponse("Failed to delete message")
		c.ServeJSON()
		return
	}
	if num < 1 {
		c.Data["json"] = helpers.CreateResponse("No messages deleted")
		c.ServeJSON()
		return
	}
	c.Data["json"] = helpers.CreateResponse("OK")
	c.ServeJSON()
}
