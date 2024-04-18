package xqkAll

import xqkVar "github.com/xqk/xqk-go-tool/xqk_var"

// XqkRuneListOpDeduplicate 去重，按顺序保留
func XqkRuneListOpDeduplicate(list *[]rune) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetDeduplicate(*list)
}

// XqkRuneListOpDeleteByIndex 根据索引删除，超出索引不处理
func XqkRuneListOpDeleteByIndex(list *[]rune, delIndex int) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetDeleteByIndex(*list, delIndex)
}

// XqkRuneListOpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkRuneListOpDeleteByRangeIndex(list *[]rune, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// XqkRuneListOpFilter 过滤切片，保留返回ture的
func XqkRuneListOpFilter(list *[]rune, keep func(rune) bool) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetFilter(*list, keep)
}

// XqkRuneListOpPop 切片元素出栈，nil、空切片都会报错
func XqkRuneListOpPop(list *[]rune) (rune, error) {
	if list == nil {
		return 0, xqkVar.BaseErrAddrNil
	}
	pop, tempList, err := XqkRuneListGetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// XqkRuneListOpReverse 切片元素反转
func XqkRuneListOpReverse(list *[]rune) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetReverse(*list)
}

// XqkRuneListOpShuffle 切片元素乱序排列
func XqkRuneListOpShuffle(list *[]rune) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetShuffle(*list)
}

// XqkRuneListOpMap 遍历计算切片，修改原切片
func XqkRuneListOpMap(list *[]rune, opFunc func(int, rune) rune) {
	if list == nil {
		return
	}
	*list = XqkRuneListGetMap(*list, opFunc)
}
