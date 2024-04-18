package xqkAll

import xqkVar "github.com/xqk/xqk-go-tool/xqk_var"

// XqkByteListOpDeduplicate 获取去重切片
func XqkByteListOpDeduplicate(list *[]byte) {
	if list == nil {
		return
	}
	*list = XqkByteListGetDeduplicate(*list)
}

// XqkByteListOpDeleteByIndex 根据索引删除，超出索引不处理
func XqkByteListOpDeleteByIndex(list *[]byte, delIndex int) {
	if list == nil {
		return
	}
	*list = XqkByteListGetDeleteByIndex(*list, delIndex)
}

// XqkByteListOpDeleteByRangeIndex 根据范围索引删除
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkByteListOpDeleteByRangeIndex(list *[]byte, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = XqkByteListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// XqkByteListOpFilter 过滤切片元素，保留返回ture的
func XqkByteListOpFilter(list *[]byte, keep func(x byte) bool) {
	if list == nil {
		return
	}
	*list = XqkByteListGetFilter(*list, keep)
}

// XqkByteListOpPop 切片元素出栈，nil、空切片都会报错
func XqkByteListOpPop(list *[]byte) (byte, error) {
	if list == nil {
		return 0, xqkVar.BaseErrAddrNil
	}
	pop, tempList, err := XqkByteListGetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// XqkByteListOpReverse 切片元素顺序反转
func XqkByteListOpReverse(list *[]byte) {
	if list == nil {
		return
	}
	*list = XqkByteListGetReverse(*list)
}

// XqkByteListOpShuffle 切片元素乱序排列
func XqkByteListOpShuffle(list *[]byte) {
	if list == nil {
		return
	}
	*list = XqkByteListGetShuffle(*list)
}

// XqkByteListOpMap 获取遍历计算后的切片，不改变原切片
func XqkByteListOpMap(list *[]byte, opFunc func(int, byte) byte) {
	if list == nil {
		return
	}
	*list = XqkByteListGetMap(*list, opFunc)
}
