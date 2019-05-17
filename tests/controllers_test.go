package test

import (
	"BeeMail/database"
	"BeeMail/helpers"
	"BeeMail/models"
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	cleanUp()
	os.Exit(code)
}

func TestGetEmptyAddresses(t *testing.T) {
	cleanUp()
	r, _ := http.NewRequest("GET", "/addresses", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Be Empty", func() {
			So(w.Body.String(), ShouldEqual, "null")
		})
	})
}

func TestGetEmptyMessages(t *testing.T) {
	cleanUp()
	r, _ := http.NewRequest("GET", "/get", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Be Empty", func() {
			So(w.Body.String(), ShouldEqual, "[]")
		})
	})
}

func TestGetAddressesAfterMessageReceived(t *testing.T) {
	cleanUp()
	addressToInsert := "10.10.10.10"
	mail := models.Mail{Type: models.Incoming, RemoteAddress: addressToInsert}
	db := *(database.GetInstance())
	_, err := db.Insert(&mail)
	r, _ := http.NewRequest("GET", "/addresses", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Database insertion should be successful", func() {
			So(err, ShouldBeNil)
		})
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.String(), ShouldEqual, fmt.Sprintf("[\"%s\"]", addressToInsert))
		})
	})
}

func cleanUp() {
	var dbPath = "./mails.db"
	dir, err := os.Getwd()
	if err != nil {
		println("Failed to get current directory for cleanup")
		return
	}
	if filepath.Base(dir) == "tests" && helpers.CheckIfFileExists(dbPath) {
		var err = os.Remove(dbPath)
		if err != nil {
			println("Failed to do cleanup after test")
		}
	}
}
