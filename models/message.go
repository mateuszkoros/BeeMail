package models

import (
	"mime/multipart"
	"strings"
)

type Message struct {
	// Id      uint   `json:"id" orm:"pk;auto"`
	// Subject string `json:"subject"`
	// Message string `json:"message"`
	Id         uint   `storm:"id,increment"`
	Subject    string `storm:"index"`
	Message    string `storm:"index"`
	Attachment multipart.File
}

func init() {
	// orm.RegisterModel(new(Message))
}

func (m Message) IsEmpty() bool {
	return strings.TrimSpace(m.Subject) == "" && strings.TrimSpace(m.Message) == ""
}
