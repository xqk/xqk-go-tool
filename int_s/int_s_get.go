package xqkSInt

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// GetMax 获取参数中最大的值
func GetMax(iList ...int) (int, error) {
	return xqkAll.XqkSIntGetMax(iList...)
}

// GetMin 获取参数中最小的值
func GetMin(iList ...int) (int, error) {
	return xqkAll.XqkSIntGetMin(iList...)
}

// GetSum 参数求和
func GetSum(iList ...int) int {
	return xqkAll.XqkSIntGetSum(iList...)
}
