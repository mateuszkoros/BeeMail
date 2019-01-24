package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"BeeMail/models"
	"github.com/astaxie/beego"
	"io"
	"os"
	"regexp"
)

type IncomingMailController struct {
	beego.Controller
}

func (c *IncomingMailController) Post() {
	var mail models.Mail
	var subject, message string
	c.Ctx.Request.ParseMultipartForm(32 << 20)
	file, handler, _ := c.Ctx.Request.FormFile("Attachment")
	if len(c.Ctx.Request.Form["Subject"]) > 0 {
		subject = c.Ctx.Request.Form["Subject"][0]
	}
	if len(c.Ctx.Request.Form["Message"]) > 0 {
		message = c.Ctx.Request.Form["Message"][0]
	}
	mail.SetSubject(subject)
	mail.SetMessage(message)
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
		c.Data["json"] = map[string]string{"Response": "Mail provided in improper format"}
	} else {
		c.Data["json"] = &mail
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
