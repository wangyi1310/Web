package model

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2"
	"os"
)

// connect DB
var (
	session *mgo.Session
	db      *mgo.Database
)

func init() {
	var err error
	session, err = mgo.Dial("0.0.0.0:27017")
	db = session.DB("mytest")
	if err != nil {
		logs.Error("connect mongo db is fail error: %v", err)
		os.Exit(1)
	}

}
