package model

//热门话题和评论信息的结构体

type RawData struct {
	Text string `json:"text"`
}

type Url struct {
	Url string `json:"url"`
}

type HotTitle struct {
	MetricsArea RawData `json:"metricsArea"`
	TitleArea   RawData `json:"titleArea"`
	ExcerptArea RawData `json:"excerptArea"` //标题背景
	Link        Url     `json:"link"`        //对应的话题URL
	Id          string
}

//定义用户信息和相关的评论

type AuthorRaw struct {
	Name      string `json:"name"`
	Url_token string `json:"url_token"`
	Work      string `json:"work"`
	Sex       string `json:"sex"`
}

type UserInfo struct {
	Data []UserRaw `json:"data"`
	Url  string    `json:"url"`
}

//mongo中存储的数据的格式如下
type UserRaw struct {
	Author  AuthorRaw `json:"author"`  //作者的信息
	Content string    `json:"content"` //评论信息
	Id      int64     `json:"id"`
}

type HotTitleCommits struct {
	HotTitles HotTitle  `json:"hot_titles"`
	UserRaws  []UserRaw `json:"user_raws"`
}

type Data struct {
	HotTitlesCommit HotTitleCommits `json:"hot_titles_commit"`
}

const (
	CommitRawUrl   = "https://www.zhihu.com/api/v4/questions/"
	BillbordRawUrl = "https://www.zhihu.com/billboard"
	OffSet         = 20 //每次读完评论之后的偏移
)
