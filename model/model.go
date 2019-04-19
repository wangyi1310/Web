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
	for _,data := range datas[0].HotTitlesCommit.UserRaws{
		comments = append(comments, data)
	}
	return  comments
}