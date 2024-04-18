package xqkByteList

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// OpDeduplicate 获取去重切片
func OpDeduplicate(list *[]byte) {
	xqkAll.XqkByteListOpDeduplicate(list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]byte, delIndex int) {
	xqkAll.XqkByteListOpDeleteByIndex(list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]byte, startIndex, endIndex int) {
	xqkAll.XqkByteListOpDeleteByRangeIndex(list, startIndex, endIndex)
}

// OpFilter 过滤切片元素，保留返回ture的
func OpFilter(list *[]byte, keep func(x byte) bool) {
	xqkAll.XqkByteListOpFilter(list, keep)
}

// OpMap 获取遍历计算后的切片，不改变原切片
func OpMap(list *[]byte, opFunc func(int, byte) byte) {
	xqkAll.XqkByteListOpMap(list, opFunc)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]byte) (byte, error) {
	return xqkAll.XqkByteListOpPop(list)
}

// OpReverse 切片元素顺序反转
func OpReverse(list *[]byte) {
	xqkAll.XqkByteListOpReverse(list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]byte) {
	xqkAll.XqkByteListOpShuffle(list)
}
