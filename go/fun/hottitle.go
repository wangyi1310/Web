package fun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//获取网页界面上的原始数据
func GetRawHotData() string{
	client :=&http.Client{}
	resp, err := client.Get(BillbordRawUrl)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil{
		fmt.Print(err)
		return ""
	}

	return string(body)
}

//查找到原始数据中热门数据相关的数据信息

func GetHotTitlefromRawData() []string {
	RawData :=GetRawHotData()
	if RawData == ""{
		return []string{}
	}

	re :=regexp.MustCompile(`"metricsArea"(.*?)}}`)
	params :=re.FindAllString(RawData,-1)


	//后期修改成lambda表达式的形式
	for index,param :=range params {
		params[index] = fmt.Sprintf("{%s",param)
	}
	return params
}

//保存热门话题的信息并转成struct
func ConvertJsontoStruct(titles []string) []HotTitle{

	var hot_titles []HotTitle
	for _,title := range titles{
		var hottitle HotTitle
		err := json.Unmarshal([]byte(title),&hottitle)
		if err != nil{
			fmt.Print(err)
			continue
		}
		//获取对应id
		hottitle.Id=GetTitleIdFromUrl(hottitle.Link.Url)
		hot_titles=append(hot_titles, hottitle)
	}
	return hot_titles
}

//通过URL获取热门话题的ID号
func GetTitleIdFromUrl(url string) string{
	re :=regexp.MustCompile(`\d{6,10}`)
	id :=re.FindString(url)
	if id == ""{
		fmt.Print("Get Question Id Fail")
	}

	return id
}



