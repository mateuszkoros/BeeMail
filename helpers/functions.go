package helpers

import (
	"github.com/astaxie/beego"
)

func CheckError(err error) {
	if err != nil {
		beego.Emergency(err)
		panic(err)
	}
}
