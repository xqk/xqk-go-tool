package xqkStr

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// ToByteList 将字符串转换为切片
//
// "Hello Xqk!" >> ['H', 'e', 'l', 'l', 'o', ' ', 'Y', 'i', 'u', '!']
//
// "" >> []
func ToByteList(str string) []byte {
	return xqkAll.XqkStrToByteList(str)
}

// ToInt 字符串>>Int，出错就返回0
func ToInt(str string) int {
	return xqkAll.XqkStrToInt(str)
}

// ToRuneList 返回Unicode代码点切片，即可以包含汉字
func ToRuneList(str string) []rune {
	return xqkAll.XqkStrToRuneList(str)
}

// ToStrList 字符串分割成切片
// "Hello Xqk!" >> ["H", "e", "l", "l", "o", " ", "Y", "i", "u", "!"]
func ToStrList(str string) []string {
	return xqkAll.XqkStrToStrList(str)
}

// ToStrListBySNum 将字符串按照数字切割，每一个数字代表输出的位数
//
// "Hello Xqk!" > [1, 2, 3] > ["H", "el", "lo "]
//
// "Hello Xqk!" > [1, 2, 3, 4] > ["H", "el", "lo ", "Xqk!"]
//
// "Hello Xqk!" > [1, 2, 3, 4, 5] > ["H", "el", "lo ", "Xqk!"]
//
// "Hello Xqk!" > [0, 1, 0, 2, 3, 4, 5] > ["", "H", "", "el", "lo ", "Xqk!"]
func ToStrListBySNum(str string, indexArr ...int) []string {
	return xqkAll.XqkStrToStrListBySNum(str, indexArr...)
}

// ToStrListBySSep 字符串通过多个字符 sep 分割成切片
func ToStrListBySSep(str string, sepList ...string) []string {
	return xqkAll.XqkStrToStrListBySSep(str, sepList...)
}

// ToStrListBySep 字符串通过指定字符 sep 分割成切片
//
// "Hope you are happy every day!" > " " > ["Hope", "you", "are", "happy", "every", "day!"]
//
// "Hope you are happy every day!" > "H" > ["", "ope you are happy every day!"]
//
// "Hope you are happy every day!" > "!" > ["Hope you are happy every day", ""]
//
// "Hello Xqk!" > "" > ["H", "e", "l", "l", "o", " ", "Y", "i", "u", "!"]
func ToStrListBySep(str, sep string) []string {
	return xqkAll.XqkStrToStrListBySep(str, sep)
}

// ToStrListNoEmptyBySSep 字符串通过多个字符 sep 分割成切片
// 该分割会将结果中所有的空字符串删除
func ToStrListNoEmptyBySSep(str string, sepList ...string) []string {
	return xqkAll.XqkStrToStrListNoEmptyBySSep(str, sepList...)
}
