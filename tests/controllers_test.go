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

var testAddress = "10.10.10.10"

func TestMain(m *testing.M) {
	code := m.Run()
	cleanUp()
	os.Exit(code)
}

func TestGetEmptyAddresses(t *testing.T) {
	r, _ := http.NewRequest("GET", "/addresses", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Getting Addresses From Empty Database\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Be Empty", func() {
			So(w.Body.String(), ShouldEqual, "null")
		})
	})
}

func TestGetEmptyMessages(t *testing.T) {
	r, _ := http.NewRequest("GET", "/get", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Getting Messages From Empty Database\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Be Empty", func() {
			So(w.Body.String(), ShouldEqual, "[]")
		})
	})
}

func TestGetAddressesAfterMessageReceived(t *testing.T) {
	insertMessageToDatabase(t)
	r, _ := http.NewRequest("GET", "/addresses", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Getting Addresses\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.String(), ShouldEqual, fmt.Sprintf("[\"%s\"]", testAddress))
		})
	})
	purgeDatabaseAfterInsertion(t)
}

func TestGetMessagesAfterMessageReceived(t *testing.T) {
	insertMessageToDatabase(t)
	r, _ := http.NewRequest("GET", fmt.Sprintf("/get?address=%s", testAddress), nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Getting Messages\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.String(), ShouldNotEqual, "[]")
		})
	})
	purgeDatabaseAfterInsertion(t)
}

func insertMessageToDatabase(t *testing.T) {
	mail := models.Mail{Type: models.Incoming, RemoteAddress: testAddress}
	db := *(database.GetInstance())
	_, err := db.Insert(&mail)
	Convey("Subject: Test Message Insertion\n", t, func() {
		So(err, ShouldBeNil)
	})
}

func purgeDatabaseAfterInsertion(t *testing.T) {
	db := *(database.GetInstance())
	_, err := db.Delete(&models.Mail{Id: 1})
	Convey("Subject: Test Message Deletion\n", t, func() {
		So(err, ShouldBeNil)
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
