package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
    beego.Router("/commitdata", &controllers.MainController{},"*:GetCommitData")
}
