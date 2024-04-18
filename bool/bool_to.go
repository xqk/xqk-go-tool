package xqkBool

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// ToInt bool转换成Int
func ToInt(b bool) int {
	return xqkAll.XqkBoolToInt(b)
}

// ToStr bool转换成string
func ToStr(b bool) string {
	return xqkAll.XqkBoolToStr(b)
}
