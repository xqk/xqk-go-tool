package xqkIntList

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// OpDeduplicate 去重，按顺序保留
func OpDeduplicate(list *[]int) {
	xqkAll.XqkIntListOpDeduplicate(list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]int, delIndex int) {
	xqkAll.XqkIntListOpDeleteByIndex(list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]int, startIndex, endIndex int) {
	xqkAll.XqkIntListOpDeleteByRangeIndex(list, startIndex, endIndex)
}

// OpFilter 过滤切片，保留返回ture的
func OpFilter(list *[]int, keep func(int) bool) {
	xqkAll.XqkIntListOpFilter(list, keep)
}

// OpMap 遍历计算切片，修改原切片
func OpMap(list *[]int, opFunc func(int, int) int) {
	xqkAll.XqkIntListOpMap(list, opFunc)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]int) (int, error) {
	return xqkAll.XqkIntListOpPop(list)
}

// OpReverse 切片元素反转
func OpReverse(list *[]int) {
	xqkAll.XqkIntListOpReverse(list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]int) {
	xqkAll.XqkIntListOpShuffle(list)
}
