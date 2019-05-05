package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"test/funcpy"
)

var (
	mutex 	sync.Mutex
	c =&http.Client{}
	Datas []Data
)

func ClearSql() error{
	err :=db.C("data").DropCollection()
	if err!=nil {
		logs.Error(err)
		return err
	}

	//clearData
	Datas = []Data{}
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
    err :=db.C("data").Find(nil).All(&Datas)
    logs.Debug("get data finish")
	SetConentStatus()
	SetUserInfo()
    return err
}

func SetConentStatus(){
	logs.Debug("set content status start")
	for i,data := range Datas{
		for j,u := range data.HotTitlesCommit.UserRaws{
			go funcpy.GetEmResult(u.Content,&Datas[i].HotTitlesCommit.UserRaws[j].Status)
			logs.Debug(Datas[i].HotTitlesCommit.UserRaws[j].Status,u.Content)
		}
	}
	logs.Debug("set content status finish")
}

func SetUserInfo(){
	for i,data := range Datas{
		for j,_ := range data.HotTitlesCommit.UserRaws{
			go GetUserInfo(&Datas[i].HotTitlesCommit.UserRaws[j].Author)
			logs.Debug(Datas[i].HotTitlesCommit.UserRaws[j].Author)
		}
	}
	logs.Debug("finish set user info")
}

func GetHotTileData() []HotTitle{
	var hots []HotTitle
	for _,data := range Datas{
		hots=append(hots,data.HotTitlesCommit.HotTitles)
	}
	return hots
}

func GetCommentData() []UserRaw{
	var comments []UserRaw
	for _,data := range Datas[0].HotTitlesCommit.UserRaws{
		comments = append(comments, data)
	}
	return  comments
}

//根据title的id获取对应的评论数据
func GetHotTileItem(title_id string) []UserRaw{
	var comments []UserRaw
	for _,data := range Datas{
		if data.HotTitlesCommit.HotTitles.Id == title_id{
			for _,data := range data.HotTitlesCommit.UserRaws{
				comments = append(comments, data)
			}
			break;
		}
	}
	return comments
}

func GetUserUrl(user_token string) string{
	return "https://www.zhihu.com/people/"+user_token+"/activities"
}

func GetUserInfoRaw(user_token string) string{
	mutex.Lock()
	defer mutex.Unlock()
	if user_token == ""{
		return ""
	}

	resp, err := c.Get(GetUserUrl(user_token))
	if err != nil{
		logs.Error(err.Error())
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
	if err != nil{
		fmt.Print(err)
		return ""
	}

	return string(body)
}

func GetUserInfo(authorRaw *AuthorRaw){
	raw :=GetUserInfoRaw(authorRaw.Url_token)
	ParseUserInfo(raw,authorRaw)
}

func ParseUserInfo(raw string,authorRaw *AuthorRaw){
	if strings.Contains(raw,"Icon--male") {
		authorRaw.Sex = "male"
	}else{
		authorRaw.Sex = "female"
	}

	re :=regexp.MustCompile(`.5.5.5.5V4H7z"/></g></svg></div>(.*?)<div`)
	str :=re.FindString(raw)
	r := []rune(str)

	cnstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])

		}
	}
	authorRaw.Work =cnstr
}