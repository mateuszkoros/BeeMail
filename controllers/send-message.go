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

type SendMessageController struct {
	beego.Controller
}

func (c *SendMessageController) Post() {
	destinationAddress := "http://localhost:1944"
	response, err := http.PostForm(destinationAddress, url.Values{
		"Subject": {"Sample subject"},
		"Message": {"Sample message"}})

	helpers.CheckError(err)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	beego.Info(string(body))

	msg := models.Message{}
	json.Unmarshal(body, &msg)
	c.Data["json"] = msg
	c.ServeJSON()
}
