package xqkAll

import (
	"strings"

	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

// XqkStrGetFirstByte 获取字符串的第一个Byte，空字符串报错
//
// "Hello Xqk!" >> 'H', nil
//
// "" >> ' ',  XqkVar.BaseErrStrEmpty
func XqkStrGetFirstByte(str string) (byte, error) {
	if str == "" {
		return ' ', xqkVar.BaseErrStrEmpty
	}
	return XqkStrToByteList(str)[0], nil
}

// XqkStrGetFirstByteNoErr 获取字符串的第一个Byte，空字符串获取空格Byte
//
// "Hello Xqk!" >> 'H'
//
// "" >> ' '
func XqkStrGetFirstByteNoErr(str string) byte {
	result, _ := XqkStrGetFirstByte(str)
	return result
}

// XqkStrGetLastByte 获取字符串最后一个Byte，空字符串报错
//
// "Hello Xqk!" >> '!', nil
//
// "" >> ' ', XqkVar.BaseErrStrEmpty
func XqkStrGetLastByte(str string) (byte, error) {
	if str == "" {
		return ' ', xqkVar.BaseErrStrEmpty
	}
	return XqkStrToByteList(str)[len(str)-1], nil
}

// XqkStrGetLastByteNoErr 获取字符串的最后一个Byte，空字符串获取空格Byte
//
// "Hello Xqk!" >> '!'
//
// "" >> ' '
func XqkStrGetLastByteNoErr(str string) byte {
	result, _ := XqkStrGetLastByte(str)
	return result
}

// XqkStrGetTrimSStr 获取去除两边指定字符串后的字符串，不修改原字符串
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
func XqkStrGetTrimSStr(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, trimByStr, targetStrArr...)
}

// XqkStrGetTrimLeftSStr 获取去除左边指定字符串后的字符串，不修改原字符串
func XqkStrGetTrimLeftSStr(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, trimLeftByStr, targetStrArr...)
}

// XqkStrGetTrimRightSStr 获取去除右边指定字符串后的字符串，不修改原字符串
func XqkStrGetTrimRightSStr(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, trimRightByStr, targetStrArr...)
}

// trimByStr 由于 strings.Trim 不能满足我的要求，过度删除字符了
// strings.Trim 好像是按第二个参数的所有rune来判断，即非第二个参数直接等于
func trimByStr(s, v string) string {
	s = trimLeftByStr(s, v)
	s = trimRightByStr(s, v)
	return s
}

func trimLeftByStr(s, v string) string {
	for {
		if strings.HasPrefix(s, v) {
			s = s[len(v):]
		} else {
			break
		}
	}
	return s
}

func trimRightByStr(s, v string) string {
	for {
		if strings.HasSuffix(s, v) {
			s = s[:strings.LastIndex(s, v)]
		} else {
			break
		}
	}
	return s
}

func trimWithoutTarget(
	str string,
	trimFunc func(s, cutset string) string,
	targetStrArr ...string) string {
	for {
		strLength := len(str)
		changeLength := 0
		for _, v := range targetStrArr {
			str = trimFunc(str, v)
			changeLength = len(str)
		}
		if strLength == changeLength {
			return str
		}
	}
}

// XqkStrGetDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果，不改变原字符串
func XqkStrGetDeleteSStr(str string, targetStrArr ...string) string {
	for _, v := range targetStrArr {
		str = strings.ReplaceAll(str, v, "")
	}
	return str
}

// XqkStrGetIndexAndSubBySStr 从多个子字符串中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果，如果两个子字符串处于相同位置，返回参数中排在前面的
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Xqk!" > ["Xqk", "ell", "ello", ""] > 1, "ell"
//
// "" > ["ell", "Xqk", "ello", ""] > -1, ""
func XqkStrGetIndexAndSubBySStr(str string, subListArr ...string) (int, string) {
	if str == "" {
		return -1, ""
	}
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		if v == "" {
			subListIndex[i] = -1
		} else {
			subListIndex[i] = strings.Index(str, v)
		}
	}
	minNotInvalidIndex := -1
	minIndexNumber := -1
	for i, v := range subListIndex {
		if v != -1 {
			if minIndexNumber == -1 || minIndexNumber > v {
				minNotInvalidIndex = i
				minIndexNumber = v
			}
		}
	}
	if minNotInvalidIndex != -1 {
		return minIndexNumber, subListArr[minNotInvalidIndex]
	}
	return -1, ""
}

// XqkStrGetIndexAndFirstSubBySStr 从多个子字符串中获取索引，返回第一个的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Xqk!" > ["Xqk", "ell", "ello", ""] > 6, "Xqk"
//
// "" > ["ell", "Xqk", "ello", ""] > -1, ""
func XqkStrGetIndexAndFirstSubBySStr(str string, subListArr ...string) (int, string) {
	if str == "" {
		return -1, ""
	}
	for _, v := range subListArr {
		tIndex := strings.Index(str, v)
		if v != "" && tIndex != -1 {
			return tIndex, v
		}
	}
	return -1, ""
}

// XqkStrGetStrByRuneIndex 根据rune长度获取字符串中的字符串
// "你好呀，Hello Xqk!" > 2 > "呀", nil
//
// "" > 2 > "", XqkError.ErrStrEmpty
//
// "1" > 20 > "呀", XqkError.ErrIndexOutOfBound
func XqkStrGetStrByRuneIndex(str string, i int) (string, error) {
	runeList := XqkStrToRuneList(str)
	if len(runeList) == 0 {
		return "", xqkVar.BaseErrStrEmpty
	}
	if i < 0 || i > len(runeList) {
		return "", xqkVar.BaseErrIndexOutOfBound
	}
	return string(runeList[i]), nil
}

// XqkStrGetFirstRune 获取第一个Rune元素
// "你好" >> 20320
// "" >> 0
// "Hello" >> 72
func XqkStrGetFirstRune(str string) rune {
	if str == "" {
		return 0
	}
	list := XqkStrToRuneList(str)
	return list[0]
}

// XqkStrGetFirstRuneIntStr 获取第一个Rune元素int的string
// "你好" >> "20320"
// "" >> "0"
// "Hello" >> "72"
func XqkStrGetFirstRuneIntStr(str string) string {
	return XqkRuneToIntStr(XqkStrGetFirstRune(str))
}

// XqkStrGetFirstRuneStr 获取第一个Rune字符串
// "你好" >> "你"
// "" >> ""
// "Hello" >> "H"
func XqkStrGetFirstRuneStr(str string) string {
	if str == "" {
		return ""
	}
	list := XqkStrToRuneList(str)
	return string(list[0])
}

// XqkStrGetLastRune 获取最后一个Rune元素
// "你好" >> 22909
// "" >> 0
// "Hello" >> 111
func XqkStrGetLastRune(str string) rune {
	if str == "" {
		return 0
	}
	list := XqkStrToRuneList(str)
	return list[len(list)-1]
}

// XqkStrGetLastRuneIntStr 获取最后一个Rune元素int的string
// "你好" >> "22909"
// "" >> "0"
// "Hello" >> "111"
func XqkStrGetLastRuneIntStr(str string) string {
	return XqkRuneToIntStr(XqkStrGetLastRune(str))
}

// XqkStrGetLastRuneStr 获取最后一个Rune字符串
// "你好" >> "好"
// "" >> ""
// "Hello" >> "o"
func XqkStrGetLastRuneStr(str string) string {
	if str == "" {
		return ""
	}
	list := XqkStrToRuneList(str)
	return string(list[len(list)-1])
}

// XqkStrGetDelEndRNStr 获取去除结尾的\r\n
func XqkStrGetDelEndRNStr(str string) string {
	return XqkStrGetTrimRightSStr(str, "\r", "\n")
}

// XqkStrGetTureStr 获取 true 的字符串
func XqkStrGetTureStr() string {
	return "true"
}

// XqkStrGetFalseStr 获取 false 的字符串
func XqkStrGetFalseStr() string {
	return "false"
}

// XqkStrGetReplaceEndStr 替换结尾的指定字符串
func XqkStrGetReplaceEndStr(s, end, r string) string {
	if strings.HasSuffix(s, end) {
		return s[:strings.LastIndex(s, end)] + r
	} else {
		return s
	}
}
