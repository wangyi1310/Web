package model

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

// connect DB
var (
	session *mgo.Session
	db      *mgo.Database
)

func init() {
	var err error
	session, err = mgo.Dial("localhost:27017")
	db = session.DB("mytest")

	if err != nil {
		fmt.Errorf("connect mongo db is fail error: %v", err)
	}

}
