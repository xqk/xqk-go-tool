package xqkAll

import xqkVar "github.com/xqk/xqk-go-tool/xqk_var"

// XqkIntListOpDeduplicate 去重，按顺序保留
func XqkIntListOpDeduplicate(list *[]int) {
	if list == nil {
		return
	}
	*list = XqkIntListGetDeduplicate(*list)
}

// XqkIntListOpDeleteByIndex 根据索引删除，超出索引不处理
func XqkIntListOpDeleteByIndex(list *[]int, delIndex int) {
	if list == nil {
		return
	}
	*list = XqkIntListGetDeleteByIndex(*list, delIndex)
}

// XqkIntListOpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkIntListOpDeleteByRangeIndex(list *[]int, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = XqkIntListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// XqkIntListOpFilter 过滤切片，保留返回ture的
func XqkIntListOpFilter(list *[]int, keep func(int) bool) {
	if list == nil {
		return
	}
	*list = XqkIntListGetFilter(*list, keep)
}

// XqkIntListOpPop 切片元素出栈，nil、空切片都会报错
func XqkIntListOpPop(list *[]int) (int, error) {
	if list == nil {
		return 0, xqkVar.BaseErrAddrNil
	}
	pop, tempList, err := XqkIntListGetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// XqkIntListOpReverse 切片元素反转
func XqkIntListOpReverse(list *[]int) {
	if list == nil {
		return
	}
	*list = XqkIntListGetReverse(*list)
}

// XqkIntListOpShuffle 切片元素乱序排列
func XqkIntListOpShuffle(list *[]int) {
	if list == nil {
		return
	}
	*list = XqkIntListGetShuffle(*list)
}

// XqkIntListOpMap 遍历计算切片，修改原切片
func XqkIntListOpMap(list *[]int, opFunc func(int, int) int) {
	if list == nil {
		return
	}
	*list = XqkIntListGetMap(*list, opFunc)
}
