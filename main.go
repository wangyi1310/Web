package main

import (
	"github.com/astaxie/beego"
	"test/fun"
	_ "test/routers"
)


func main() {
	go fun.TimeRun()
	beego.Run()
}


