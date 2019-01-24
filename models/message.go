package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Message struct {
	Id      uint   `json:"id" orm:"pk;auto"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func init() {
	orm.RegisterModel(new(Message))
}

func (m Message) IsEmpty() bool {
	return strings.TrimSpace(m.Subject) == "" && strings.TrimSpace(m.Message) == ""
}

func (m *Message) SetSubject(subject string) {
	trimmedSubject := strings.TrimSpace(subject)
	if trimmedSubject != "" {
		m.Subject = trimmedSubject
	}
}

func (m *Message) SetMessage(message string) {
	trimmedMessage := strings.TrimSpace(message)
	if trimmedMessage != "" {
		m.Message = trimmedMessage
	}
}
