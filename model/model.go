package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)


var datas []Data
func ClearSql() error{
	err :=db.C("data").DropCollection()
	if err!=nil {
		logs.Error(err)
		return err
	}

	//clearData
	datas = []Data{}

	return nil
}


func InsertHotTitle(data Data) error{
	err :=db.C("data").Insert(data)
	if err!=nil{
		fmt.Errorf(err.Error())
		return err
	}
	return nil
}


func GetData() error{
	return db.C("data").Find(nil).All(&datas)
}

func GetHotTileData() []HotTitle{
	var hots []HotTitle
	for _,data := range datas{
		hots=append(hots,data.HotTitlesCommit.HotTitles)
	}
	return hots
}

func GetCommentData() []UserRaw{
	var comments []UserRaw
	for _,data := range datas[5].HotTitlesCommit.UserRaws{
		comments = append(comments, data)
	}
	return  comments
}

//根据title的id获取对应的评论数据
func GetHotTileItem(title_id string) []UserRaw{
	var comments []UserRaw
	logs.Warn(title_id,len(datas))
	for _,data := range datas{
		logs.Warn(data.HotTitlesCommit.HotTitles.Id)
		if data.HotTitlesCommit.HotTitles.Id == title_id{
			for _,data := range data.HotTitlesCommit.UserRaws{
				comments = append(comments, data)
			}
			break;
		}
	}
	return comments
}
