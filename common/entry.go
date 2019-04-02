package common

type DefaultOutMsg struct {
	Code int         `json:code`
	Msg  string      `json:msg`
	Data interface{} `json: data`
}
