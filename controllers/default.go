package controllers

import (
	"github.com/astaxie/beego"
	"test/common"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	
	var hot_themes []common.HotTheme
	for a:=0;a<20;a++ {
		hot_themes = append(hot_themes,
			common.HotTheme{
				ThemeName:     "test",
				ThemePopulors: 0,
				ThemeUrl:      "",
			})
	}
	c.Data["Themes"] = common.DefaultOutMsg{
		Length: len(hot_themes),
		Data: hot_themes,
	}
	c.TplName = "index.html"
}


func (c *MainController) GetHotTitleData(){
	

}