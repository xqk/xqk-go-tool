package xqkStrList

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// ToStr 将切片合并输出
func ToStr(list []string) string {
	return xqkAll.XqkStrListToStr(list)
}

// ToStrBySep 按照分隔符将切片输出
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
func ToStrBySep(list []string, sep string) string {
	return xqkAll.XqkStrListToStrBySep(list, sep)
}
