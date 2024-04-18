package xqkByte

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// ToStr 单个byte转字符串，无论是byte数组、切片都可以调用 string(b) 转换
func ToStr(b byte) string {
	return xqkAll.XqkByteToStr(b)
}
