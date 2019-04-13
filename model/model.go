package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

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