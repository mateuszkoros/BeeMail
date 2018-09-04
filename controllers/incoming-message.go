package controllers

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"BeeMail/models"
	"github.com/astaxie/beego"
	"io"
	"os"
	"regexp"
	"strings"
)

type IncomingMessageController struct {
	beego.Controller
}

func (c *IncomingMessageController) Post() {
	var msg models.Message
	var subject, message string
	c.Ctx.Request.ParseMultipartForm(32 << 20)
	file, handler, _ := c.Ctx.Request.FormFile("Attachment")
	if len(c.Ctx.Request.Form["Subject"]) > 0 {
		subject = c.Ctx.Request.Form["Subject"][0]
	}
	if len(c.Ctx.Request.Form["Message"]) > 0 {
		message = c.Ctx.Request.Form["Message"][0]
	}
	if strings.TrimSpace(subject) != "" {
		msg.Subject = subject
	}
	if strings.TrimSpace(message) != "" {
		msg.Message = message
	}
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

// remove potentially harmful characters from filename
func validateFileName(fileName string) string {
	validator, err := regexp.Compile(`[*\\/"\[\]:;|=,&]`)
	if err != nil {
		helpers.CheckError(err)
	}
	return validator.ReplaceAllString(fileName, "")
}
