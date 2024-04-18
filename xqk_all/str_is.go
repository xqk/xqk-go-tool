package xqkAll

import (
	"strings"
)

// XqkStrIsEmpty 判断是否为空，空格不为空
//
// ""    >> true
//
// "   " >> false
func XqkStrIsEmpty(str string) bool {
	return str == ""
}

// XqkStrIsNotEmpty 判断是否不为空
//
// ""    >> false
//
// "   " >> true
func XqkStrIsNotEmpty(str string) bool {
	return !XqkStrIsEmpty(str)
}

// XqkStrIsBlank 判断是否全是空格或者空
//
// ""    >> true
//
// "   " >> true
func XqkStrIsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

// XqkStrIsNotBlank 判断是否全部非空格或者空
func XqkStrIsNotBlank(str string) bool {
	return !XqkStrIsBlank(str)
}

// XqkStrIsLetter 判断是否是字母a-zA-Z，如果是字母返回true
//
// "" >> false
//
// "A" >> true
//
// "你" >> false
func XqkStrIsLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	return XqkByteIsLetter(XqkStrToByteList(str)[0])
}

func XqkStrIsLowerLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	return XqkByteIsLowerLetter(XqkStrToByteList(str)[0])
}

func XqkStrIsUpperLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	return XqkByteIsUpperLetter(XqkStrToByteList(str)[0])
}

// XqkStrIsNotLetter 判断是否不是字母a-zA-Z，如果不是字母返回true
func XqkStrIsNotLetter(str string) bool {
	return !XqkStrIsLetter(str)
}

// XqkStrIsStartWithLetterByte 是否以字母Byte开头
func XqkStrIsStartWithLetterByte(str string) bool {
	b, err := XqkStrGetFirstByte(str)
	if err != nil {
		return false
	}
	return XqkByteIsLetter(b)
}

// XqkStrIsEndWithLetterByte 是否以字母Byte结尾
func XqkStrIsEndWithLetterByte(str string) bool {
	b, err := XqkStrGetLastByte(str)
	if err != nil {
		return false
	}
	return XqkByteIsLetter(b)
}

// XqkStrIsStartWithLetterRune 是否以字母Rune开头
func XqkStrIsStartWithLetterRune(str string) bool {
	return XqkStrIsLetter(XqkStrGetFirstRuneStr(str))
}

// XqkStrIsEndWithLetterRune 是否以字母Rune结尾
func XqkStrIsEndWithLetterRune(str string) bool {
	return XqkStrIsLetter(XqkStrGetLastRuneStr(str))
}

// XqkStrIsTrue 判读字符串是否是true
//
// "TRUE" >> true
//
// "True" >> true
//
// "TrUe" >> true
//
// "true" >> true
//
// "1" >> true
//
// "anyStr" >> false
//
// "not1" >> false
func XqkStrIsTrue(str string) bool {
	return strings.ToLower(str) == strings.ToLower("true") || str == "1"
}

// XqkStrIsFalse 判读字符串是否是false，结果和 IsTure 相反
//
// "TRUE" >> false
//
// "True" >> false
//
// "TrUe" >> false
//
// "true" >> false
//
// "1" >> false
//
// "anyStr" >> true
//
// "not1" >> true
func XqkStrIsFalse(str string) bool {
	return !XqkStrIsTrue(str)
}

// XqkStrIsContainsAnyByte 是否包含对比串的任意Byte
func XqkStrIsContainsAnyByte(s, chars string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(chars); j++ {
			if s[i] == chars[j] {
				return true
			}
		}
	}
	return false
}

// XqkStrIsContainsAnyRune 是否包含对比串的任意Rune
func XqkStrIsContainsAnyRune(s, chars string) bool {
	sRune := XqkStrToRuneList(s)
	charsRune := XqkStrToRuneList(chars)
	for i := 0; i < len(sRune); i++ {
		for j := 0; j < len(charsRune); j++ {
			if sRune[i] == charsRune[j] {
				return true
			}
		}
	}
	return false
}

// XqkStrIsGt 字符串按照unicode逐一比较，是否大于
func XqkStrIsGt(s1, s2 string) bool {
	r1 := XqkStrToRuneList(s1)
	r2 := XqkStrToRuneList(s2)
	return XqkRuneListIsGtByUnicode(r1, r2)
}

// XqkStrIsGe 字符串按照unicode逐一比较，是否大于等于
func XqkStrIsGe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return XqkStrIsGt(s1, s2)
}

// XqkStrIsLt 字符串按照unicode逐一比较，是否小于
func XqkStrIsLt(s1, s2 string) bool {
	return !XqkStrIsGt(s1, s2)
}

// XqkStrIsLe 字符串按照unicode逐一比较，是否小于等于
func XqkStrIsLe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return XqkStrIsLt(s1, s2)
}
