package xqkSByteList

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// GetMaxLengthEl 获取长度最大的第一个切片
func GetMaxLengthEl(list ...[]byte) []byte {
	return xqkAll.XqkSByteListGetMaxLengthEl(list...)
}

// GetMerge 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func GetMerge(bListArr ...[]byte) []byte {
	return xqkAll.XqkSByteListGetMerge(bListArr...)
}

// GetMinLengthEl 获取长度最小的第一个切片
func GetMinLengthEl(list ...[]byte) []byte {
	return xqkAll.XqkSByteListGetMinLengthEl(list...)
}
