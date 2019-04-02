package controllers

import (
	"github.com/astaxie/beego"
	"test/common"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}


func (c *MainController) GetHotTitleData(){
	c.Data["json"] = common.DefaultOutMsg{
		Code:0,
		Msg:"success",
	}
	c.ServeJSON()

}