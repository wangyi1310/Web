package fun

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"test/model"
	"time"
)

var smutex sync.Mutex
var wg2 sync.WaitGroup
func GetUserInfoRaw(user_token string) string {
	smutex.Lock()
	defer smutex.Unlock()
	if user_token == "" {
		return ""
	}
	c := http.Client{}
	resp, err := c.Get(GetUserUrl(user_token))

	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	return string(body)
}

func GetUserInfo(authorRaw model.AuthorRaw) {
	raw := GetUserInfoRaw(authorRaw.Url_token)
	ParseUserInfo(raw, authorRaw)
	wg2.Done()
}

func ParseUserInfo(raw string, authorRaw model.AuthorRaw)  {
	if strings.Contains(raw, "Icon--male") {
		authorRaw.Sex = "male"
	} else {
		authorRaw.Sex = "female"
	}

	re := regexp.MustCompile(`.5.5.5.5V4H7z"/></g></svg></div>(.*?)<div`)
	str := re.FindString(raw)
	r := []rune(str)

	cnstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
		}
	}
	authorRaw.Work = cnstr
	logs.Error(authorRaw)
	model.InsertUserInfo(authorRaw)
}

func GetUserUrl(user_token string) string {
	return "https://www.zhihu.com/people/" + user_token + "/activities"
}

func SetUserInfo() {
	logs.Info("SetUserInfo")
	data := model.GetData()
	for _, d := range data {
		for _, u := range d.HotTitlesCommit.UserRaws {
			if u.Author.Url_token == "" {
				continue
			}
			wg2.Add(1)
		    go GetUserInfo(u.Author)
		}
	}
	wg2.Wait()
	logs.Info("FinishInfo")
	time.Sleep(time.Second * 5000 * 2)
}