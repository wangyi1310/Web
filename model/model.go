package model

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

var (
	mutex  sync.Mutex
	Datas []Data
)


func ClearSql() error {
	err := db.C("data").DropCollection()
	if err != nil {
		logs.Error(err)
		return err
	}
	Datas = []Data{}

	return nil
}

func InsertHotTitle(data Data) error {
	err := db.C("data").Insert(&data)
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

func InsertUserInfo(userInof AuthorRaw) error {
	err := db.C("userinfo").Insert(userInof)
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

func GetData() []Data {
	Datas = []Data{}
	db.C("data").Find(nil).All(&Datas)
	logs.Debug("get data finish")
	return Datas
}

func GetHotTileData() []HotTitle {
	logs.Info("get hot data")
	var hots []HotTitle
	Datas := GetData()
	for _, data := range Datas {
		hots = append(hots, data.HotTitlesCommit.HotTitles)
	}
	logs.Info("finsh hot data")
	return hots
}

//根据title的id获取对应的评论数据
func GetHotTileItem(title_id string) []UserRaw {
	var comments []UserRaw
	Datas := GetData()
	for _, data := range Datas {
		if data.HotTitlesCommit.HotTitles.Id == title_id {
			for _, data := range data.HotTitlesCommit.UserRaws {
				comments = append(comments, data)
			}
			break
		}
	}
	return comments
}


func EmotionClassData(id string) (int, int, int) {
	Datas := GetData()
	logs.Debug("set content status start")
	var (
		PosCount int
		NegCount int
		NoCount  int
	)

	for _, data := range Datas {
		if data.HotTitlesCommit.HotTitles.Id == id || id == "000000"{
			for _, u := range data.HotTitlesCommit.UserRaws {
				if u.Status == "积极" {
					PosCount++
				} else if u.Status == "消极" {
					NegCount++
				} else {
					NoCount++
				}
			}
		}

	}
	return PosCount, NegCount, NoCount
}

func SexClassCount(id string)([]int,error){
	malePosCount,err :=db.C("userinfo").Find(bson.M{"sex": "male","id":id,"status" : "积极"}).Count()
	if err != nil {
		logs.Error( "get sex male count failuer error: %v", err)
		return []int{},err
	}

	maleNegCount,err :=db.C("userinfo").Find(bson.M{"sex": "male","id":id,"status" : "消极"}).Count()
	if err != nil {
		logs.Error( "get sex male count failuer error: %v", err)
		return []int{},err
	}

	maleNoCount,err :=db.C("userinfo").Find(bson.M{"sex": "male","id":id,"status" : "中立"}).Count()
	if err != nil {
		logs.Error( "get sex male count failuer error: %v", err)
		return []int{},err
	}

	feMalePosCount,err :=db.C("userinfo").Find(bson.M{"sex": "female","id":id,"status" : "积极"}).Count()
	if err != nil {
		logs.Error( "get sex male count failuer error: %v", err)
		return []int{},err
	}

	feMaleNegCount,err :=db.C("userinfo").Find(bson.M{"sex": "female","id":id,"status" : "消极"}).Count()
	if err != nil {
		logs.Error( "get sex male count failuer error: %v", err)
		return []int{},err
	}

	feMaleNoCount,err :=db.C("userinfo").Find(bson.M{"sex": "female","id":id,"status" : "中立"}).Count()
	if err != nil {
		logs.Error( "get sex male count failuer error: %v", err)
		return []int{},err
	}

	return []int{malePosCount,maleNegCount,maleNoCount,feMalePosCount,feMaleNegCount,feMaleNoCount},nil
}

func SeachUser(username string) AuthorRaw{
	var user AuthorRaw
	err:=db.C("userinfo").Find(bson.M{"name": username}).One(&user)
	if err !=nil{
		return AuthorRaw{}
	}
	return user
}