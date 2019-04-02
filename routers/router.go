package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
	beego.Router("/hottitle", &controllers.MainController{}, "Get:GetHotTitleData")
}
