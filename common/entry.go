package common

type DefaultOutMsg struct {
	Data   interface{}
	Length int
}


//定义热门话题的结构体
type HotTheme struct {
	ThemeName     string
	ThemePopulors string
	ThemeUrl      string
	ThemeIndex    int
}

type Comment struct {
	Content string
	Status  string
	Url     string
	ThemeIndex    int
}
