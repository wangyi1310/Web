package common

func CutString(str string,l int) string{
	strRune := []rune(str)
	if len(strRune) <= l{
		return str
	}
	str =string((strRune[:l]))
	str = str +"..."
	return str
}

func SortArray(){

}