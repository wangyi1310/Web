package routers

import (
	"github.com/astaxie/beego"
	"test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/comment", &controllers.MainController{}, "*:GetCommentData")
	beego.Router("/comments", &controllers.MainController{}, "*:GetCommentsDatas")
	beego.Router("/emotionclass", &controllers.MainController{}, "*:EmotionClass")
	beego.Router("/sexclass", &controllers.MainController{}, "*:SexClass")
}
