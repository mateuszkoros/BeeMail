package helpers

import (
	"BeeMail/models"
	"github.com/astaxie/beego"
	"net/http"
	"os"
)

func CheckError(err error) {
	if err != nil {
		beego.Emergency(err)
		panic(err)
	}
}

func CreateMailFromHttpRequest(request *http.Request) models.Mail {
	m := models.Mail{}
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

func CheckIfFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic("Cannot determine whether file exists")
	}
}
