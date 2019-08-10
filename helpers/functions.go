package helpers

import (
	"BeeMail/models"
	"github.com/astaxie/beego"
	"net"
	"net/http"
	"os"
	"strings"
)

// CheckError is a helper function that logs a fatal error and exits application.
func CheckError(err error) {
	if err != nil {
		beego.Emergency(err)
		panic(err)
	}
}

// CreateMailFromHttpRequest creates Message object from HTTP request's
// Subject, Message, AttachmentName and Attachment field.
func CreateMailFromHttpRequest(request *http.Request) models.Mail {
	m := models.Mail{}
	if len(request.Form["Subject"]) > 0 {
		m.SetSubject(request.Form["Subject"][0])
	}
	if len(request.Form["Message"]) > 0 {
		m.SetMessage(request.Form["Message"][0])
	}
	if len(request.Form["AttachmentName"]) > 0 {
		m.SetAttachmentName(request.Form["AttachmentName"][0])
	}
	if len(request.Form["Attachment"]) > 0 {
		m.SetAttachment(request.Form["Attachment"][0])
	}
	return m
}

// CreateResponse is a helper function for creating HTTP response containing specified message.
func CreateResponse(text string) *models.ReceiverResponse {
	return &models.ReceiverResponse{Response: text}
}

// CheckIfFileExists checks whether file exists on disk.
// It causes application to panic if it cannot unequivocally determine file's existence.
func CheckIfFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic("Cannot determine whether file exists")
	}
}

// CheckIfLocalAddress checks if request came from localhost.
// It is used by endpoints accepting only local requests.
func CheckIfLocalAddress(address string) bool {
	addressArray := strings.Split(address, ":")
	address = strings.Join(addressArray[:len(addressArray)-1], ":")
	parsedIp := net.ParseIP(address)
	return parsedIp.IsLoopback()
}
