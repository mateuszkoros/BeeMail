package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"github.com/astaxie/beego"
	"io"
	"os"
	"regexp"
)

type IncomingMailController struct {
	beego.Controller
}

func (c *IncomingMailController) Post() {
	mail := helpers.CreateMailFromHttpRequest(c.Ctx.Request)
	c.Ctx.Request.ParseMultipartForm(32 << 20)
	file, handler, _ := c.Ctx.Request.FormFile("Attachment")
	if file != nil {
		defer file.Close()
		filename := validateFileName(handler.Filename)
		f, err := os.Create("./" + filename)
		if err != nil {
			beego.Warn(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
	if mail.IsEmpty() {
		c.Data["json"] = helpers.CreateResponse("Mail provided in improper format")
	} else {
		c.Data["json"] = helpers.CreateResponse("Message received")
		db := *(database.GetInstance())
		db.Insert(&mail)
	}
	beego.Info("Received mail:\n" + string(c.Ctx.Input.RequestBody))
	c.ServeJSON()
}

// remove potentially harmful characters from filename
func validateFileName(fileName string) string {
	validator, err := regexp.Compile(`[*\\/"\[\]:;|=,&]`)
	if err != nil {
		helpers.CheckError(err)
	}
	return validator.ReplaceAllString(fileName, "")
}
