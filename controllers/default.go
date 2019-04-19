package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"test/common"
	"test/model"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {

	err := model.GetData()
	if err != nil{
		logs.Error("get hot data errro")
	}
	hs := model.GetHotTileData()
    cs := model.GetCommentData()
	var hot_themes []common.HotTheme
    var hot_comments [] common.Comment
	for a,h := range hs{
		if a +1 == 30{
			break;
		}
		hot_themes = append(hot_themes,
			common.HotTheme{
				ThemeName:     common.CutString(h.TitleArea.Text,17),
				ThemePopulors: h.MetricsArea.Text,
				ThemeUrl:      "",
				ThemeIndex:    a+1,
			})
	}

	for a,c := range cs{
		if a +1 == 18{
			break;
		}
		hot_comments = append(hot_comments,
			common.Comment{
			Content: 	common.CutString(c.Content,115),
			Url:        c.Author.Name,
			ThemeIndex: a+1,
			Status: "积极",
			},
		)
	}

	c.Data["Themes"] = common.DefaultOutMsg{
		Length: len(hot_themes),
		Data: hot_themes,
	}
	c.Data["Comments"] = common.DefaultOutMsg{
		Length: len(hot_comments),
		Data: hot_comments,
	}
	c.TplName = "index.html"
}


func (c *MainController) GetHotTitleData(){


}