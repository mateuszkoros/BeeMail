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

// TODO add sent messages to database
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
		helpers.CheckError(err)

		receiverResponse := getResponseData(response)
		responses = append(responses, receiverResponse)

		c.Data["json"] = receiverResponse
		err = response.Body.Close()
		helpers.CheckError(err)
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
