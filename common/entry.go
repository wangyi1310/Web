package common

type DefaultOutMsg struct {
	Data   interface{}
	Length int
}


//定义热门话题的结构体
type HotTheme struct {
	ThemeName     string
	ThemePopulors int
	ThemeUrl      string
}
