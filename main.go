package main

import (
	"github.com/astaxie/beego"
	"test/go/fun"
	_ "test/routers"
)

func main() {
	beego.Run()
	go fun.GetCommits()
}

