package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"test/common"
	"test/fun"
	"test/model"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	hs := model.GetHotTileData()
	var hot_themes []common.HotTheme
	for a, h := range hs {
		hot_themes = append(hot_themes,
			common.HotTheme{
				ThemeName:     common.CutString(h.TitleArea.Text, 17),
				ThemePopulors: common.GetNumByString(h.MetricsArea.Text),
				ThemeIndex:    a + 1,
				ThemeId:       h.Id,
			})
	}
	hot_themes = common.Sorts(hot_themes)
	var length = len(hot_themes)
	if length >29{
		length =29
	}
	c.Data["Themes"] = common.DefaultOutMsg{
		Length: length,
		Data:   hot_themes[:length],
		Msg:    0,
	}
	c.TplName = "index.html"
}


func (c *MainController) GetCommentData() {
	title_id := c.GetString("title_id")
	if title_id == "" {
		logs.Error("title id is null")
	}
	var hot_comments []common.Comment
	cs := model.GetHotTileItem(title_id)
	for a, c := range cs {
		if a+1 == 17 {
			break
		}
		hot_comments = append(hot_comments,
			common.Comment{
				Content:    common.CutString(c.Content, 115),
				UserInfo:   c.Author.Name,
				ThemeIndex: a + 1,
				Status:    c.Status,
				Url: fun.GetUserUrl(c.Author.Url_token),
			},
		)
	}
	c.Data["json"] = common.DefaultOutMsg{hot_comments, len(hot_comments), 0}
	c.ServeJSON()
	return
}

func(c *MainController) EmotionClass(){
	title_id := c.GetString("title_id")
	if title_id == "" {
		logs.Error("title id is null")
	}
    pos,neg,noEmo := model.EmotionClassData(title_id)
	c.Data["json"] = common.DefaultOutMsg{common.Emotion{pos,neg,noEmo}, 0, 0}
	c.ServeJSON()
	return
}

func(c *MainController) SexClass(){
	title_id := c.GetString("title_id")
	if title_id == "" {
		logs.Error("title id is null")
	}
	results,err :=model.SexClassCount()
	if err!=nil{
		c.Data["json"] = common.DefaultOutMsg{0, 0, 1}
		c.ServeJSON()
		return
	}

	c.Data["json"] = common.DefaultOutMsg{common.SexCount{
		Male:results[0],
		Female: results[1],
	}, 0, 0}
	c.ServeJSON()
}


func (c *MainController) GetCommentsDatas() {
	title_id := c.GetString("title_id")
	if title_id == "" {
		logs.Error("title id is null")
	}
	var hot_comments []common.Comment
	cs := model.GetHotTileItem(title_id)
	for a, c := range cs {
		hot_comments = append(hot_comments,
			common.Comment{
				Content:    common.CutString(c.Content, 115),
				UserInfo:   c.Author.Name,
				ThemeIndex: a + 1,
				Status:    c.Status,
				Url: fun.GetUserUrl(c.Author.Url_token),
			},
		)
	}
	c.Data["json"] = common.DefaultOutMsg{hot_comments, len(hot_comments), 0}
	c.ServeJSON()
	return
}