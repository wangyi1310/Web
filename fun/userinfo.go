package fun

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"test/model"
)

func GetUserUrl(user_token string) string{
	return "https://www.zhihu.com/people/"+user_token+"/activities"
}

func GetUserInfoRaw(user_token string) string{
	if user_token == ""{
		logs.Error("user token is null")
		return ""
	}
	client :=&http.Client{}
	resp, err := client.Get(GetUserUrl(user_token))
	if err != nil{
		logs.Error(err.Error())
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil{
		fmt.Print(err)
		return ""
	}

	return string(body)
}


func ParseUserInfo(raw string,authorRaw *model.AuthorRaw){
	if strings.Contains(raw,"Icon--male") {
		authorRaw.Sex = "male"
	}else{
		authorRaw.Sex = "female"
	}

	re :=regexp.MustCompile("2.002H15zM7 4V2.5s-.004-.5.5-.5h5c.5 0 .5.5.5.5V4H7z/></g></svg></div>(.*?)<div")
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

func GetUserInfo(authorRaw *model.AuthorRaw){
	raw :=GetUserInfoRaw(authorRaw.Url_token)
	ParseUserInfo(raw,authorRaw)
}

