package xqkAll

// XqkIntListToStr 将所有int转成字符串并输出
func XqkIntListToStr(list []int) string {
	return XqkIntListToStrBySep(list, "")
}

// XqkIntListToStrBySep 按照分隔符将切片输出
//
// [1, 2] > " " > "1 2"
//
// [123, 654, 963] > "-" > "123-654-963"
//
// [] > "-" > ""
//
// [1] > "=" > "1"
//
// [1, 2] > "" > "12"
func XqkIntListToStrBySep(list []int, sep string) string {
	if XqkIntListIsEmpty(list) {
		return ""
	}
	result := ""
	for i := range list {
		if i == 0 {
			result += XqkIntToStr(list[i])
		} else {
			result += sep + XqkIntToStr(list[i])
		}
	}
	return result
}
