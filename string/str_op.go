package xqkStr

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// OpDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果
func OpDeleteSStr(str *string, targetStrArr ...string) {
	xqkAll.XqkStrOpDeleteSStr(str, targetStrArr...)
}

// OpFormatPathSeparator 格式化字符串中的文件分隔符为当前系统的文件分隔符
func OpFormatPathSeparator(str *string) {
	xqkAll.XqkStrOpFormatPathSeparator(str)
}

// OpReplaceEndStr 替换结尾指定字符串
func OpReplaceEndStr(str *string, end, r string) {
	xqkAll.XqkStrOpReplaceEndStr(str, end, r)
}

// OpTrimLeftSStr 获取去除左边指定字符串后的字符串
func OpTrimLeftSStr(str *string, targetStrArr ...string) {
	xqkAll.XqkStrOpTrimLeftSStr(str, targetStrArr...)
}

// OpTrimRightSStr 获取去除右边指定字符串后的字符串
func OpTrimRightSStr(str *string, targetStrArr ...string) {
	xqkAll.XqkStrOpTrimRightSStr(str, targetStrArr...)
}

// OpTrimSStr 获取去除两边指定字符串后的字符串
//
// "  --1Hello Xqk!+ " > [" ", "--", "+", "1"] > "Hello Xqk!"
//
// "Hello Xqk!Hello Xqk!" > ["Hello", "--", "+", "!"] > " Xqk!Hello Xqk"
//
// " \n\r\n Hello Xqk!  \n\r\n  " > ["\n", "\r", " "] > "Hello Xqk!"
func OpTrimSStr(str *string, targetStrArr ...string) {
	xqkAll.XqkStrOpTrimSStr(str, targetStrArr...)
}
