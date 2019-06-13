package common

type DefaultOutMsg struct {
	Data   interface{} `json:"data"`
	Length int         `json:"length"`
	Msg    int         `json:"msg"`
}

//定义热门话题的结构体
type HotTheme struct {
	ThemeName     string
	ThemePopulors int
	ThemeIndex    int
	ThemeId       string
}
type newlist []HotTheme

type Comment struct {
	Sex        string `json:"sex"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	UserInfo   string `json:"user_info"`
	ThemeIndex int    `json:"theme_index"`
	Url       string  `json:"url"`
	Work      string `json:"work"`
}

type Emotion struct {
	Pos   int `json:"pos"`
	Neg   int `json:"neg"`
	NoEmo int `json:"no_emo"`

	MalePosCount   int `json:"male_pos_count"`
	MaleNegCount   int `json:"male_neg_count"`
	MaleNoCount    int `json:"male_no_count"`

	FemalePosCount   int `json:"female_pos_count"`
	FemaleNegCount   int `json:"female_neg_count"`
	FemaleNoCount    int `json:"female_no_count"`
}

type SexCount struct {
	Male int `json:"male"`
	Female int `json:"female"`
}