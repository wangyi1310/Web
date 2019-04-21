package common

import (
	"regexp"
	"strconv"
)

func CutString(str string,l int) string{
	strRune := []rune(str)
	if len(strRune) <= l{
		return str
	}
	str =string((strRune[:l]))
	str = str +"..."
	return str
}

func GetNumByString(str string)int{
	re := regexp.MustCompile("(\\w*[0-9]+)\\w*")
	s :=re.FindString(str)
	num,_:=strconv.Atoi(s)
	return num
}

func Sorts(I newlist)newlist{
	for i := 0; i < len(I)-1; i++ {
		for j := 0; j < len(I)-1-i; j++ {
			if I[j].ThemePopulors < I[j+1].ThemePopulors{
				temp := I[j]
				I[j].ThemePopulors=I[j+1].ThemePopulors
				I[j].ThemeId=I[j+1].ThemeId
				I[j].ThemeName=I[j+1].ThemeName
				I[j+1].ThemePopulors = temp.ThemePopulors
				I[j+1].ThemeId = temp.ThemeId
				I[j+1].ThemeName = temp.ThemeName
			}
		}
	}
	return I
}