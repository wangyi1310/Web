package engine



//请求数据结构
type Request struct {
	Url string    //请求url
	ParserFunc func(interface{}) ParseResult    //内容解析函数采用函数指针的形式
}

//经过内容解析函数解析后的返回数据结构体
type ParseResult struct {
	Items interface{}       //抓取到的有用信息项
}