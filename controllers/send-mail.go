package controllers

import (
	"BeeMail/helpers"
	"BeeMail/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SendMailController struct {
	beego.Controller
}

func (c *SendMailController) Post() {
	destinationAddress := "http://localhost:1944"
	response, err := http.PostForm(destinationAddress, url.Values{
		"Subject": {"Sample subject"},
		"Message": {"Sample message"}})

	helpers.CheckError(err)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	beego.Info(string(body))

	mail := models.Mail{}
	json.Unmarshal(body, &mail)
	c.Data["json"] = mail
	c.ServeJSON()
}
