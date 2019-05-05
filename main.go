package main

import (
	"github.com/astaxie/beego"
	"test/model"
	_ "test/routers"
)

func main() {
	//go fun.TimeRun()
	go model.GetData()
	beego.Run()
}



