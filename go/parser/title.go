package parser

import (
	"awesomeProject/engine"
	"awesomeProject/fun"
	"awesomeProject/model"
	"fmt"
	"regexp"
)

//解析话题的相关数据



func ParseTitleData(RawData string) *engine.ParseResult {

	if RawData == ""{
		return nil
	}


	re :=regexp.MustCompile(`"metricsArea"(.*?)}}`)
	params :=re.FindAllString(RawData,-1)

	//后期修改成lambda表达式的形式
	for index,param :=range params {
		params[index] = fmt.Sprintf("{%s",param)
	}

	result := engine.ParseResult{
		Items:fun.ConvertJsontoStruct(params),
	}

	return &result
}

//解析评论数据
func ParseCommitData(RawData model.UserInfo) *engine.ParseResult{
	result := engine.ParseResult{
		Items: RawData,
	}

	return &result
}
