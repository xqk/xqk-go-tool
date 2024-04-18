package xqkAll

import (
	"sort"

	"github.com/psampaz/slice"
	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

// XqkStrListGetDeduplicate 获取去重切片，不改变原切片
func XqkStrListGetDeduplicate(list []string) []string {
	return slice.DeduplicateString(list)
}

// XqkStrListGetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func XqkStrListGetDeleteByIndex(list []string, delIndex int) []string {
	result, err := slice.DeleteString(list, delIndex)
	if err != nil {
		return list
	}
	return result
}

// XqkStrListGetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkStrListGetDeleteByRangeIndex(list []string, startIndex, endIndex int) []string {
	result, err := slice.DeleteRangeString(list, startIndex, endIndex)
	if err != nil {
		return list
	}
	return result
}

// XqkStrListGetFilter 过滤切片元素，保留返回ture的，不改变原切片
func XqkStrListGetFilter(list []string, keep func(x string) bool) []string {
	return slice.FilterString(list, keep)
}

// XqkStrListGetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func XqkStrListGetPop(list []string) (string, []string, error) {
	popString, i, err := slice.PopString(list)
	if err != nil {
		return "", nil, xqkVar.BaseErrListEmpty
	}
	return popString, i, err
}

// XqkStrListGetReverse 切片元素顺序反转，不改变原切片
func XqkStrListGetReverse(list []string) []string {
	return slice.ReverseString(list)
}

// XqkStrListGetShuffle 切片元素乱序排列，不改变原切片
func XqkStrListGetShuffle(list []string) []string {
	return slice.ShuffleString(list)
}

// XqkStrListGetMap 获取遍历计算后的切片，不改变原切片
func XqkStrListGetMap(list []string, opFunc func(int, string) string) []string {
	if XqkStrListIsEmpty(list) {
		return list
	}
	outList := make([]string, len(list), len(list))
	for i, v := range list {
		outList[i] = opFunc(i, v)
	}
	return outList
}

// XqkStrListGetCopy 拷贝一个切片
func XqkStrListGetCopy(list []string) []string {
	return slice.CopyString(list)
}

// XqkStrListGetIndexByEl 获取元素索引，如果没有该元素则返回-1
func XqkStrListGetIndexByEl(list []string, el string) int {
	if XqkStrListIsEmpty(list) {
		return -1
	}
	for i, v := range list {
		if v == el {
			return i
		}
	}
	return -1
}

// XqkStrListGetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
//
// ["a", "b", "c", "d"] > ["b", "c"] > 1
//
// ["a", "b", "c", "d"] > ["y", "c"] > -1
//
// [] > ["y", "c"] > -1
//
// nil > ["y", "c"] > -1
//
// ["a", "b"] > [] > -1
//
// ["a", "b"] > nil > -1
//
// [] > [] > -1
//
// nil > nil > -1
func XqkStrListGetIndexByList(list []string, subList []string) int {
	if XqkStrListIsEmpty(list) {
		return -1
	}
	n := len(subList)
	switch {
	case n == 0:
		return -1
	case n == 1:
		return XqkStrListGetIndexByEl(list, subList[0])
	case n == len(list):
		if XqkStrListIsEqual(list, subList) {
			return 0
		} else {
			return -1
		}
	case n > len(list):
		return -1
	}
	for i := 0; i < len(list)-n; i++ {
		if XqkStrListIsEqual(list[i:i+n], subList) {
			return i
		}
	}
	return -1
}

// XqkStrListGetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// ["a", "b", "c", "d", "e", "f", "g"] > [{"b", "c"},{},{"d", "e"}] > 1
func XqkStrListGetIndexBySList(list []string, subListArr ...[]string) int {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = XqkStrListGetIndexByList(list, v)
	}
	XqkIntListOpFilter(&subListIndex, func(i int) bool {
		return i != -1
	})
	min, err := XqkIntListGetMin(subListIndex)
	if err != nil {
		return -1
	}
	return min
}

// XqkStrListGetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// ["a", "b", "c", "d", "e", "f", "g"] > [{},{"b", "c"},{"b", "c", "d"},{"e", "f", "g"}] > 1, ["b", "c"]
func XqkStrListGetIndexAndSubBySList(list []string, subListArr ...[]string) (int, []string) {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = XqkStrListGetIndexByList(list, v)
	}
	minNotInvalidIndex := -1
	minIndexNumber := -1
	for i, v := range subListIndex {
		if v != -1 {
			if minIndexNumber == -1 || minIndexNumber < v {
				minNotInvalidIndex = i
				minIndexNumber = v
			}
		}
	}
	if minNotInvalidIndex != -1 {
		return minIndexNumber, subListArr[minNotInvalidIndex]
	}
	return -1, nil
}

// XqkStrListGetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func XqkStrListGetElByIndex(list []string, index int) string {
	if XqkStrListIsEmpty(list) || index < 0 || index >= len(list) {
		return ""
	}
	return list[index]
}

// XqkStrListGetDeleteEmptyEl 去除所有空串，不改变原切片
//
// ["", "Hello", "", "Xqk"] >> ["Hello", "Xqk"]
//
// [" ", "Hello", "", "Xqk"] >> [" ", "Hello", "Xqk"]
//
// [] >> []
//
// nil >> nil
func XqkStrListGetDeleteEmptyEl(list []string) []string {
	return XqkStrListGetFilter(list, func(x string) bool {
		return !XqkStrIsEmpty(x)
	})
}

// XqkStrListGetDeleteBlankEl 去除所有空串和空格字符串，不改变原切片
//
// ["", "Hello", "", "Xqk"] >> ["Hello", "Xqk"]
//
// [" ", "Hello", "", "Xqk"] >> ["Hello", "Xqk"]
//
// [] >> []
//
// nil >> nil
func XqkStrListGetDeleteBlankEl(list []string) []string {
	return XqkStrListGetFilter(list, func(x string) bool {
		return !XqkStrIsBlank(x)
	})
}

// XqkStrListGetAscUnicode 按照Unicode逐个升序排列
func XqkStrListGetAscUnicode(list []string) []string {
	if XqkStrListIsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i] == list[j] {
			return true
		}
		tli := list[i]
		tlj := list[j]
		return XqkRuneListIsLeByUnicode(XqkStrToRuneList(tli), XqkStrToRuneList(tlj))
	})
	return list
}

// XqkStrListGetDescUnicode 按照Unicode逐个降序排列
func XqkStrListGetDescUnicode(list []string) []string {
	if XqkStrListIsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i] == list[j] {
			return true
		}
		tli := list[i]
		tlj := list[j]
		return XqkRuneListIsGeByUnicode(XqkStrToRuneList(tli), XqkStrToRuneList(tlj))
	})
	return list
}

// XqkStrListGetAscUnicodeAndLowerBeforeUpper 按照Unicode逐个升序排列，(大小写颠倒)
func XqkStrListGetAscUnicodeAndLowerBeforeUpper(list []string) []string {
	if XqkStrListIsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i] == list[j] {
			return true
		}
		return XqkRuneListIsLeByUnicodeAndLowerBeforeUpper(XqkStrToRuneList(list[i]), XqkStrToRuneList(list[i]))
	})
	return list
}

// XqkStrListGetDescUnicodeAndLowerBeforeUpper 按照Unicode逐个降序排列，(大小写颠倒)
func XqkStrListGetDescUnicodeAndLowerBeforeUpper(list []string) []string {
	if XqkStrListIsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i] == list[j] {
			return true
		}
		return XqkRuneListIsGeByUnicodeAndLowerBeforeUpper(XqkStrToRuneList(list[i]), XqkStrToRuneList(list[i]))
	})
	return list
}
