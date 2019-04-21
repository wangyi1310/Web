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
	if err != nil {
		logs.Error("get hot data errro")
	}
	hs := model.GetHotTileData()
	cs := model.GetCommentData()
	var hot_themes []common.HotTheme
	var hot_comments []common.Comment
	for a, h := range hs {
		if a+1 == 30 {
			break
		}
		hot_themes = append(hot_themes,
			common.HotTheme{
				ThemeName:     common.CutString(h.TitleArea.Text, 17),
				ThemePopulors: common.GetNumByString(h.MetricsArea.Text),
				ThemeIndex:    a + 1,
				ThemeId:       h.Id,
			})
	}

	hot_themes = common.Sorts(hot_themes)
	for a, c := range cs {
		if a+1 == 18 {
			break
		}
		hot_comments = append(hot_comments,
			common.Comment{
				Content:    common.CutString(c.Content, 115),
				UserInfo:   c.Author.Name,
				ThemeIndex: a + 1,
				Status:     "积极",
			},
		)
	}

	c.Data["Themes"] = common.DefaultOutMsg{
		Length: len(hot_themes),
		Data:   hot_themes,
		Msg:    0,
	}
	c.Data["Comments"] = common.DefaultOutMsg{
		Length: len(hot_comments),
		Data:   hot_comments,
		Msg:    0,
	}
	c.TplName = "index.html"
}

func (c *MainController) GetHotTitleData() {

}

func (c *MainController) GetCommitData() {
	title_id := c.GetString("title_id")
	if title_id == "" {
		logs.Error("title id is null")
	}
	var hot_comments []common.Comment
	cs := model.GetHotTileItem(title_id)
	for a, c := range cs {
		if a+1 == 18 {
			break
		}
		hot_comments = append(hot_comments,
			common.Comment{
				Content:    common.CutString(c.Content, 115),
				UserInfo:   c.Author.Name,
				ThemeIndex: a + 1,
				Status:     "积极",
			},
		)
	}
	c.Data["json"] = common.DefaultOutMsg{hot_comments, len(hot_comments), 0}

	c.ServeJSON()
	return
}
