package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type MailDestination string

const (
	Outgoing = "Outgoing"
	Incoming = "Incoming"
)

type Mail struct {
	Id            uint            `json:"id" orm:"pk;auto"`
	Subject       string          `json:"subject"`
	Message       string          `json:"message"`
	Destination   MailDestination `json:"destination"`
	RemoteAddress string          `json:"remoteaddress"`
	Timestamp     time.Time       `json:"timestamp" orm:"auto_now_add;type(datetime)"`
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
	trimmedAddress := strings.TrimSpace(address)
	trimmedAddress = strings.Split(trimmedAddress, ":")[0]
	if trimmedAddress != "" {
		m.RemoteAddress = trimmedAddress
	}
}
