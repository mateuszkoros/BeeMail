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
