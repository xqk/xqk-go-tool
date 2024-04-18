package xqkByteList

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// ToStr 将元素合并成字符串
func ToStr(list []byte) string {
	return xqkAll.XqkByteListToStr(list)
}

// ToStrBySep 按照分隔符将切片输出
//
// ['a', 'b'] > " " > "a b"
//
// ['1', '6', '9'] > "-" > "1-6-9"
//
// [] > "-" > ""
//
// ['H'] > "=" > "H"
//
// ['a', 'b'] > "" > "ab"
//
// nil > "=" > ""
func ToStrBySep(list []byte, sep string) string {
	return xqkAll.XqkByteListToStrBySep(list, sep)
}
