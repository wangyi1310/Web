package fetcher

import (
	"awesomeProject/fun"
	"awesomeProject/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

//定义采集器 传入url 进行数据的采集
var client =&http.Client{}


//采集热门话题数据
func GetRawHotData(url string) string{
	resp, err := client.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil{
		fmt.Print(err)
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		return ""
	}

	return string(body)
}

//采集评论数据

func GetRawComment(url string) model.UserInfo{
	return fun.GetComment(url)
}