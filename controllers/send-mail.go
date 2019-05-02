package controllers

import (
	"BeeMail/database"
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
	if len(c.Ctx.Request.Form["Destination"]) == 0 {
		c.Data["json"] = helpers.CreateResponse("Please specify destination address")
		c.ServeJSON()
		return
	}
	mail := helpers.CreateMailFromHttpRequest(c.Ctx.Request)
	if mail.IsEmpty() {
		c.Data["json"] = helpers.CreateResponse("Mail provided in improper format")
		c.ServeJSON()
		return
	}
	var responses []models.ReceiverResponse
	for _, destination := range c.Ctx.Request.Form["Destination"] {
		response, err := http.PostForm(destination, url.Values{
			"Subject": {mail.Subject},
			"Message": {mail.Message}})
		if err != nil {
			beego.Error("Failed to send message", err)
			c.Data["json"] = helpers.CreateResponse("Failed to send message - " + err.Error())
			c.ServeJSON()
			return
		}

		receiverResponse := getResponseData(response)
		responses = append(responses, receiverResponse)

		c.Data["json"] = receiverResponse
		err = response.Body.Close()
		helpers.CheckError(err)

		mail.Type = models.Outgoing
		mail.SetRemoteAddress(destination)
		db := *(database.GetInstance())
		db.Insert(&mail)
	}
	c.ServeJSON()
}

func getResponseData(response *http.Response) models.ReceiverResponse {
	body, err := ioutil.ReadAll(response.Body)
	helpers.CheckError(err)

	receiverResponse := models.ReceiverResponse{}
	err = json.Unmarshal(body, &receiverResponse)
	helpers.CheckError(err)
	beego.Info("Received response: " + receiverResponse.Response)
	return receiverResponse
}
