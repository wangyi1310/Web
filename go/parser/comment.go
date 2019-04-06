package parser

import (
	"awesomeProject/engine"
	"awesomeProject/model"
	"encoding/json"
	"fmt"
)

//解析出每条链接中的相关评论


func ParseCommentData(RawData string) *engine.ParseResult {
	var user model.UserInfo

	err := json.Unmarshal([]byte(RawData),&user)
	if err != nil{
		fmt.Print(err)
		return nil
	}

	result := engine.ParseResult{
		Items:user,
	}

	return &result
}
