package routers

import (
	"github.com/astaxie/beego"
	"test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/commitdata", &controllers.MainController{}, "*:GetCommitData")
	beego.Router("/emotionclass", &controllers.MainController{}, "*:EmotionClass")
	beego.Router("/sexclass", &controllers.MainController{}, "*:SexClass")
}
