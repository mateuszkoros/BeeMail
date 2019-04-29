package helpers

import (
	"BeeMail/models"
	"github.com/astaxie/beego"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		beego.Emergency(err)
		panic(err)
	}
}

func CreateMailFromHttpRequest(request *http.Request) models.Mail {
	m := models.Mail{}
	m.SetRemoteAddress(request.RemoteAddr)
	// if there is no remote address in request return empty mail, signalling that request is invalid
	if m.RemoteAddress == "" {
		return m
	}
	if len(request.Form["Subject"]) > 0 {
		m.SetSubject(request.Form["Subject"][0])
	}
	if len(request.Form["Message"]) > 0 {
		m.SetMessage(request.Form["Message"][0])
	}
	return m
}

func CreateResponse(text string) *models.ReceiverResponse {
	return &models.ReceiverResponse{Response: text}
}
