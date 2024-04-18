package xqkAll

import (
	"github.com/psampaz/slice"
	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

// XqkIntListGetDeduplicate 获取去重切片，不改变原切片
func XqkIntListGetDeduplicate(list []int) []int {
	return slice.DeduplicateInt(list)
}

// XqkIntListGetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func XqkIntListGetDeleteByIndex(list []int, delIndex int) []int {
	result, err := slice.DeleteInt(list, delIndex)
	if err != nil {
		return list
	}
	return result
}

// XqkIntListGetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkIntListGetDeleteByRangeIndex(list []int, startIndex, endIndex int) []int {
	result, err := slice.DeleteRangeInt(list, startIndex, endIndex)
	if err != nil {
		return list
	}
	return result
}

// XqkIntListGetFilter 过滤切片元素，保留返回ture的，不改变原切片
func XqkIntListGetFilter(list []int, keep func(x int) bool) []int {
	return slice.FilterInt(list, keep)
}

// XqkIntListGetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func XqkIntListGetPop(list []int) (int, []int, error) {
	popString, i, err := slice.PopInt(list)
	if err != nil {
		return 0, nil, xqkVar.BaseErrListEmpty
	}
	return popString, i, err
}

// XqkIntListGetReverse 切片元素顺序反转，不改变原切片
func XqkIntListGetReverse(list []int) []int {
	return slice.ReverseInt(list)
}

// XqkIntListGetShuffle 切片元素乱序排列，不改变原切片
func XqkIntListGetShuffle(list []int) []int {
	return slice.ShuffleInt(list)
}

// XqkIntListGetMap 获取遍历计算后的切片，不改变原切片
func XqkIntListGetMap(list []int, opFunc func(int, int) int) []int {
	if XqkIntListIsEmpty(list) {
		return list
	}
	outList := make([]int, len(list), len(list))
	for i, v := range list {
		outList[i] = opFunc(i, v)
	}
	return outList
}

// XqkIntListGetCopy 拷贝一个切片
func XqkIntListGetCopy(list []int) []int {
	return slice.CopyInt(list)
}

// XqkIntListGetIndexByEl 获取元素索引，如果没有该元素则返回-1
func XqkIntListGetIndexByEl(list []int, el int) int {
	if XqkIntListIsEmpty(list) {
		return -1
	}
	for i, v := range list {
		if v == el {
			return i
		}
	}
	return -1
}

// XqkIntListGetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
//
// [1, 2, 3, 4] > [2, 3] > 1
//
// [1, 2, 3, 4] > [5, 3] > -1
//
// [] > [2, 3] > -1
//
// nil > [2, 3] > -1
//
// [1, 2] > [] > -1
//
// [1, 2] > nil > -1
//
// [] > [] > -1
//
// nil > nil > -1
func XqkIntListGetIndexByList(list []int, subList []int) int {
	if XqkIntListIsEmpty(list) {
		return -1
	}
	n := len(subList)
	switch {
	case n == 0:
		return -1
	case n == 1:
		return XqkIntListGetIndexByEl(list, subList[0])
	case n == len(list):
		if XqkIntListIsEqual(list, subList) {
			return 0
		} else {
			return -1
		}
	case n > len(list):
		return -1
	}
	for i := 0; i < len(list)-n; i++ {
		if XqkIntListIsEqual(list[i:i+n], subList) {
			return i
		}
	}
	return -1
}

// XqkIntListGetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// [1, 2, 3, 4, 5, 6, 7] > [{2, 3},{},{5, 6}] > 1
func XqkIntListGetIndexBySList(list []int, subListArr ...[]int) int {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = XqkIntListGetIndexByList(list, v)
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

// XqkIntListGetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// [1, 2, 3, 4, 5, 6, 7] > [{},{2, 3},{2, 3, 4},{5, 6, 7}] > 1, [2, 3]
func XqkIntListGetIndexAndSubBySList(list []int, subListArr ...[]int) (int, []int) {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = XqkIntListGetIndexByList(list, v)
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

// XqkIntListGetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func XqkIntListGetElByIndex(list []int, index int) int {
	if XqkIntListIsEmpty(list) || index < 0 || index >= len(list) {
		return 0
	}
	return list[index]
}

// XqkIntListGetMax 获取切片中最大的byte，nil、空切片都会报错
func XqkIntListGetMax(list []int) (int, error) {
	maxByte, err := slice.MaxInt(list)
	if err != nil {
		return 0, xqkVar.BaseErrListEmpty
	}
	return maxByte, nil
}

// XqkIntListGetMin 获取切片中最小的byte，nil、空切片都会报错
func XqkIntListGetMin(list []int) (int, error) {
	minByte, err := slice.MinInt(list)
	if err != nil {
		return 0, xqkVar.BaseErrListEmpty
	}
	return minByte, nil
}

// XqkIntListGetSum 所有int值算术相加，nil、空切片返回0
func XqkIntListGetSum(list []int) int {
	sumInt, err := slice.SumInt(list)
	if err != nil {
		return 0
	}
	return sumInt
}
