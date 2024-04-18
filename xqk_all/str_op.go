package xqkAll

import (
	"os"
	"strings"
)

// XqkStrOpTrimSStr 获取去除两边指定字符串后的字符串
//
// "  --1Hello Xqk!+ " > [" ", "--", "+", "1"] > "Hello Xqk!"
//
// "Hello Xqk!Hello Xqk!" > ["Hello", "--", "+", "!"] > " Xqk!Hello Xqk"
//
// " \n\r\n Hello Xqk!  \n\r\n  " > ["\n", "\r", " "] > "Hello Xqk!"
func XqkStrOpTrimSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = XqkStrGetTrimSStr(*str, targetStrArr...)
}

// XqkStrOpTrimLeftSStr 获取去除左边指定字符串后的字符串
func XqkStrOpTrimLeftSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = XqkStrGetTrimLeftSStr(*str, targetStrArr...)
}

// XqkStrOpTrimRightSStr 获取去除右边指定字符串后的字符串
func XqkStrOpTrimRightSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = XqkStrGetTrimRightSStr(*str, targetStrArr...)
}

// XqkStrOpDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果
func XqkStrOpDeleteSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = XqkStrGetDeleteSStr(*str, targetStrArr...)
}

// XqkStrOpFormatPathSeparator 格式化字符串中的文件分隔符为当前系统的文件分隔符
func XqkStrOpFormatPathSeparator(str *string) {
	*str = strings.ReplaceAll(*str, "/", string(os.PathSeparator))
	*str = strings.ReplaceAll(*str, "\\", string(os.PathSeparator))
}

// XqkStrOpReplaceEndStr 替换结尾指定字符串
func XqkStrOpReplaceEndStr(str *string, end, r string) {
	if str == nil {
		return
	}
	*str = XqkStrGetReplaceEndStr(*str, end, r)
}
