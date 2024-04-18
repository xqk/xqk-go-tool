package xqkAll

// XqkStrListToStr 将切片合并输出
func XqkStrListToStr(list []string) string {
	return XqkStrListToStrBySep(list, "")
}

// XqkStrListToStrBySep 按照分隔符将切片输出
//
// ["Hello", "Xqk"] > " " > "Hello Xqk"
//
// ["123", "654", "963"] > "-" > "123-654-963"
//
// [] > "-" > ""
//
// ["Hello"] > "=" > "Hello"
//
// ["Hello", "Xqk"] > "" > "HelloXqk"
func XqkStrListToStrBySep(list []string, sep string) string {
	outStr := ""
	if XqkStrListIsEmpty(list) {
		return ""
	}
	if len(list) == 1 {
		return list[0]
	}
	for i, v := range list {
		if i == 0 {
			outStr += v
			continue
		}
		outStr += sep + v
	}
	return outStr
}
