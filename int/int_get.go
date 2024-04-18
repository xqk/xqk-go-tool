package xqkInt

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// GetEnglishNumTail 获取英文的数字后缀
func GetEnglishNumTail(i int) string {
	return xqkAll.XqkIntGetEnglishNumTail(i)
}

// GetFalseInt 获取 false 的int值
func GetFalseInt() int {
	return xqkAll.XqkIntGetFalseInt()
}

// GetTrueInt 获取 true 的int值
func GetTrueInt() int {
	return xqkAll.XqkIntGetTrueInt()
}
