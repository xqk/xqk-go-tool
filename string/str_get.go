package xqkStr

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// GetDelEndRNStr 获取去除结尾的\r\n
func GetDelEndRNStr(str string) string {
	return xqkAll.XqkStrGetDelEndRNStr(str)
}

// GetDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果，不改变原字符串
func GetDeleteSStr(str string, targetStrArr ...string) string {
	return xqkAll.XqkStrGetDeleteSStr(str, targetStrArr...)
}

// GetFalseStr 获取 false 的字符串
func GetFalseStr() string {
	return xqkAll.XqkStrGetFalseStr()
}

// GetFirstByte 获取字符串的第一个Byte，空字符串报错
//
// "Hello Xqk!" >> 'H', nil
//
// "" >> ' ',  XqkVar.BaseErrStrEmpty
func GetFirstByte(str string) (byte, error) {
	return xqkAll.XqkStrGetFirstByte(str)
}

// GetFirstByteNoErr 获取字符串的第一个Byte，空字符串获取空格Byte
//
// "Hello Xqk!" >> 'H'
//
// "" >> ' '
func GetFirstByteNoErr(str string) byte {
	return xqkAll.XqkStrGetFirstByteNoErr(str)
}

// GetFirstRune 获取第一个Rune元素
// "你好" >> 20320
// "" >> 0
// "Hello" >> 72
func GetFirstRune(str string) rune {
	return xqkAll.XqkStrGetFirstRune(str)
}

// GetFirstRuneIntStr 获取第一个Rune元素int的string
// "你好" >> "20320"
// "" >> "0"
// "Hello" >> "72"
func GetFirstRuneIntStr(str string) string {
	return xqkAll.XqkStrGetFirstRuneIntStr(str)
}

// GetFirstRuneStr 获取第一个Rune字符串
// "你好" >> "你"
// "" >> ""
// "Hello" >> "H"
func GetFirstRuneStr(str string) string {
	return xqkAll.XqkStrGetFirstRuneStr(str)
}

// GetIndexAndFirstSubBySStr 从多个子字符串中获取索引，返回第一个的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Xqk!" > ["Xqk", "ell", "ello", ""] > 6, "Xqk"
//
// "" > ["ell", "Xqk", "ello", ""] > -1, ""
func GetIndexAndFirstSubBySStr(str string, subListArr ...string) (int, string) {
	return xqkAll.XqkStrGetIndexAndFirstSubBySStr(str, subListArr...)
}

// GetIndexAndSubBySStr 从多个子字符串中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果，如果两个子字符串处于相同位置，返回参数中排在前面的
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Xqk!" > ["Xqk", "ell", "ello", ""] > 1, "ell"
//
// "" > ["ell", "Xqk", "ello", ""] > -1, ""
func GetIndexAndSubBySStr(str string, subListArr ...string) (int, string) {
	return xqkAll.XqkStrGetIndexAndSubBySStr(str, subListArr...)
}

// GetLastByte 获取字符串最后一个Byte，空字符串报错
//
// "Hello Xqk!" >> '!', nil
//
// "" >> ' ', XqkVar.BaseErrStrEmpty
func GetLastByte(str string) (byte, error) {
	return xqkAll.XqkStrGetLastByte(str)
}

// GetLastByteNoErr 获取字符串的最后一个Byte，空字符串获取空格Byte
//
// "Hello Xqk!" >> '!'
//
// "" >> ' '
func GetLastByteNoErr(str string) byte {
	return xqkAll.XqkStrGetLastByteNoErr(str)
}

// GetLastRune 获取最后一个Rune元素
// "你好" >> 22909
// "" >> 0
// "Hello" >> 111
func GetLastRune(str string) rune {
	return xqkAll.XqkStrGetLastRune(str)
}

// GetLastRuneIntStr 获取最后一个Rune元素int的string
// "你好" >> "22909"
// "" >> "0"
// "Hello" >> "111"
func GetLastRuneIntStr(str string) string {
	return xqkAll.XqkStrGetLastRuneIntStr(str)
}

// GetLastRuneStr 获取最后一个Rune字符串
// "你好" >> "好"
// "" >> ""
// "Hello" >> "o"
func GetLastRuneStr(str string) string {
	return xqkAll.XqkStrGetLastRuneStr(str)
}

// GetReplaceEndStr 替换结尾的指定字符串
func GetReplaceEndStr(s, end, r string) string {
	return xqkAll.XqkStrGetReplaceEndStr(s, end, r)
}

// GetStrByRuneIndex 根据rune长度获取字符串中的字符串
// "你好呀，Hello Xqk!" > 2 > "呀", nil
//
// "" > 2 > "", XqkError.ErrStrEmpty
//
// "1" > 20 > "呀", XqkError.ErrIndexOutOfBound
func GetStrByRuneIndex(str string, i int) (string, error) {
	return xqkAll.XqkStrGetStrByRuneIndex(str, i)
}

// GetTrimLeftSStr 获取去除左边指定字符串后的字符串，不修改原字符串
func GetTrimLeftSStr(str string, targetStrArr ...string) string {
	return xqkAll.XqkStrGetTrimLeftSStr(str, targetStrArr...)
}

// GetTrimRightSStr 获取去除右边指定字符串后的字符串，不修改原字符串
func GetTrimRightSStr(str string, targetStrArr ...string) string {
	return xqkAll.XqkStrGetTrimRightSStr(str, targetStrArr...)
}

// GetTrimSStr 获取去除两边指定字符串后的字符串，不修改原字符串
//
// "  --1Hello Xqk!+ " > [" ", "--", "+", "1"] > "Hello Xqk!"
//
// "Hello Xqk!Hello Xqk!" > ["Hello", "--", "+", "!"] > " Xqk!Hello Xqk"
//
// " \n\r\n Hello Xqk!  \n\r\n  " > ["\n", "\r", " "] > "Hello Xqk!"
//
// 先删除前面的，后删除后面的，如果前面删除导致后面匹配不到是不会删除后面的，比如一下案例：
// "HHHHHH" > "HHHH" > "HH"
// 此处先删除4个H，后面2个H无法匹配
func GetTrimSStr(str string, targetStrArr ...string) string {
	return xqkAll.XqkStrGetTrimSStr(str, targetStrArr...)
}

// GetTureStr 获取 true 的字符串
func GetTureStr() string {
	return xqkAll.XqkStrGetTureStr()
}
