package fun

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"regexp"
	"test/model"
	"time"
)

//获取每条热门话题下的评论数据

//定义相关的抓取页面的相关配置数据


//组装相应的获取评论的参数


func PackGetData(offset int) map[string]string{
	data := map[string]string{
		"include":"data%5B%2A%5D.is_normal%2Cadmin_closed_comment%2Creward_info%2Cis_collapsed%2Cannotation_action%2Cannotation_detail%2Ccollapse_reason%2Cis_sticky%2Ccollapsed_by%2Csuggest_edit%2Ccomment_count%2Ccan_comment%2Ccontent%2Ceditable_content%2Cvoteup_count%2Creshipment_settings%2Ccomment_permission%2Ccreated_time%2Cupdated_time%2Creview_info%2Crelevant_info%2Cquestion%2Cexcerpt%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%2Cis_labeled%3Bdata%5B%2A%5D.mark_infos%5B%2A%5D.url%3Bdata%5B%2A%5D.author.follower_count%2Cbadge%5B%2A%5D.topics",
		"limit": "20",
		"offset": fmt.Sprintf("%d",offset),
		"platform": "desktop",
		"sort_by": "default",
	}
	return data
}


//获取评论的URL
func GetCommentUrl(url string, offest int) string{
	if url==""{
		fmt.Print("url is no null")
		return ""
	}
	params :=PackGetData(offest)
	url=url+"?"+"include="+params["include"]+"&"+"limit="+params["limit"]+"&"+"offset="+params["offset"]+"&"+"platform="+params["platform"]+"&"+"sort_by="+params["sort_by"]

	return url
}


//
func GetCommentRaw(url string) *model.UserInfo{

	client :=&http.Client{}
	resp,err := client.Get(url)
	if err !=nil{
		fmt.Print(err)
		return nil
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
	}

	var user model.UserInfo

	err = json.Unmarshal(b,&user)
	if err != nil{
		fmt.Print(err)
		return nil
	}
	return &user

}

//获取评论下的相关数据
func GetComment(url string) model.UserInfo{
	var offset int = 0
	var users model.UserInfo
	logs.Debug("start get comment")
	for{
		url := GetCommentUrl(url,offset)
		user := GetCommentRaw(url)
		if len(user.Data) != model.OffSet {
			break
		}

		for _,u := range user.Data {
			u.Content=FilterIllegalWorld(u.Content)
			users.Data=append(users.Data, u)
		}

		offset +=model.OffSet
	}
	logs.Debug("finish get comment")
	return users
}

//过滤一些html相关标签
func FilterIllegalWorld(commnent string) string{
	re :=regexp.MustCompile(`<(.*?)>|\d{6,10}`)
	commnent=re.ReplaceAllString(commnent,"")
	return commnent

}


func InsertHotTitleCommit(title model.HotTitle,url string,id int){
	userInfo :=GetComment(url)
	logs.Warn("title name %s title commit num %d",title.TitleArea,len(userInfo.Data))
	s :=model.HotTitleCommits{
		title,
		userInfo.Data,
	}
	if err :=model.InsertHotTitle(model.Data{s});err!=nil{
		logs.Error("insert value failure")
		return
	}
	//logs.Info("insert value success",id)
}
func GetCommits(){
		//清理数据库中存在的老数据
		s := ConvertJsontoStruct(GetHotTitlefromRawData())

		for i, t := range s {
			logs.Debug(i)
			//获取到每条评论后写入sql中 对应的是s然后sql中的每个话题名字对应下方的评论 每5分钟系统跑一次。
			//userInfo :=GetComment(model.CommitRawUrl + s[i].Id + "/answers")
			go InsertHotTitleCommit(t,model.CommitRawUrl + s[i].Id + "/answers",i)
		}
}

//system 10分钟清理一次sql重新拉取数据一次
func TimeRun(){
	for{
		logs.Info("start clear sql")
		model.ClearSql()
		logs.Info("finish clear sql")
		GetCommits()
		time.Sleep(time.Second*300*2)
		//每隔５分钟数据库数据重新拉去一次
	}
}

