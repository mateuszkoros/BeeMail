package models

import (
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	"time"
)

type MailType string

const (
	Outgoing = "Outgoing"
	Incoming = "Incoming"
)

type Mail struct {
	Id             uint      `json:"id" orm:"pk;auto"`
	Subject        string    `json:"subject"`
	Message        string    `json:"message"`
	Type           MailType  `json:"type"`
	RemoteAddress  string    `json:"remoteaddress"`
	AttachmentName string    `json:"attachmentname"`
	Attachment     string    `json:"attachment"`
	Timestamp      time.Time `json:"timestamp" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Mail))
}

func (m Mail) IsEmpty() bool {
	return strings.TrimSpace(m.Subject) == "" && strings.TrimSpace(m.Message) == ""
}

func (m *Mail) SetSubject(subject string) {
	trimmedSubject := strings.TrimSpace(subject)
	if trimmedSubject != "" {
		m.Subject = trimmedSubject
	}
}

func (m *Mail) SetMessage(message string) {
	trimmedMessage := strings.TrimSpace(message)
	if trimmedMessage != "" {
		m.Message = trimmedMessage
	}
}

func (m *Mail) SetRemoteAddress(address string) {
	validator, err := regexp.Compile(`^(.*://)|(:.*)$`)
	if err != nil {
		return
	}
	trimmedAddress := validator.ReplaceAllString(address, "")
	trimmedAddress = strings.TrimSpace(trimmedAddress)
	if trimmedAddress != "" {
		m.RemoteAddress = trimmedAddress
	}
}
