package xqkAll

import (
	"math/rand"
	"time"

	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

// XqkRuneListGetDeduplicate 获取去重切片，不改变原切片
func XqkRuneListGetDeduplicate(list []rune) []rune {
	if len(list) < 2 {
		return list
	}

	seen := make(map[rune]struct{})

	j := 0
	for k := range list {
		if _, ok := seen[list[k]]; ok {
			continue
		}
		seen[list[k]] = struct{}{}
		list[j] = list[k]
		j++
	}

	return list[:j]
}

// XqkRuneListGetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func XqkRuneListGetDeleteByIndex(list []rune, delIndex int) []rune {
	if len(list) == 0 {
		return list
	}
	if delIndex < 0 || delIndex > len(list)-1 {
		return list
	}
	return append(list[:delIndex], list[delIndex+1:]...)
}

// XqkRuneListGetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkRuneListGetDeleteByRangeIndex(list []rune, startIndex, endIndex int) []rune {
	if len(list) == 0 {
		return list
	}
	if startIndex < 0 || startIndex > len(list)-1 {
		return list
	}
	if endIndex < 0 || endIndex > len(list)-1 {
		return list
	}
	if startIndex > endIndex {
		return list
	}
	return append(list[:startIndex], list[endIndex+1:]...)
}

// XqkRuneListGetFilter 过滤切片元素，保留返回ture的，不改变原切片
func XqkRuneListGetFilter(list []rune, keep func(x rune) bool) []rune {
	if len(list) == 0 {
		return list
	}
	n := 0
	for _, v := range list {
		if keep(v) {
			list[n] = v
			n++
		}
	}
	return list[:n]
}

// XqkRuneListGetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func XqkRuneListGetPop(list []rune) (rune, []rune, error) {
	if len(list) == 0 {
		return 0, nil, xqkVar.BaseErrListEmpty
	}
	return list[len(list)-1], list[:len(list)-1], nil
}

// XqkRuneListGetReverse 切片元素顺序反转，不改变原切片
func XqkRuneListGetReverse(list []rune) []rune {
	if len(list) == 0 {
		return list
	}
	s, e := 0, len(list)-1
	for s < e {
		list[s], list[e] = list[e], list[s]
		s++
		e--
	}
	return list
}

// XqkRuneListGetShuffle 切片元素乱序排列，不改变原切片
func XqkRuneListGetShuffle(list []rune) []rune {
	if len(list) <= 1 {
		return list
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

// XqkRuneListGetMap 获取遍历计算后的切片，不改变原切片
func XqkRuneListGetMap(list []rune, opFunc func(int, rune) rune) []rune {
	if XqkRuneListIsEmpty(list) {
		return list
	}
	outList := make([]rune, len(list), len(list))
	for i, v := range list {
		outList[i] = opFunc(i, v)
	}
	return outList
}

// XqkRuneListGetCopy 拷贝一个切片
func XqkRuneListGetCopy(list []rune) []rune {
	return append(list[:0:0], list...)
}

// XqkRuneListGetIndexByEl 获取元素索引，如果没有该元素则返回-1
func XqkRuneListGetIndexByEl(list []rune, el rune) int {
	if XqkRuneListIsEmpty(list) {
		return -1
	}
	for i, v := range list {
		if v == el {
			return i
		}
	}
	return -1
}

// XqkRuneListGetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
func XqkRuneListGetIndexByList(list []rune, subList []rune) int {
	if XqkRuneListIsEmpty(list) {
		return -1
	}
	n := len(subList)
	switch {
	case n == 0:
		return -1
	case n == 1:
		return XqkRuneListGetIndexByEl(list, subList[0])
	case n == len(list):
		if XqkRuneListIsEqual(list, subList) {
			return 0
		} else {
			return -1
		}
	case n > len(list):
		return -1
	}
	for i := 0; i < len(list)-n; i++ {
		if XqkRuneListIsEqual(list[i:i+n], subList) {
			return i
		}
	}
	return -1
}

// XqkRuneListGetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
func XqkRuneListGetIndexBySList(list []rune, subListArr ...[]rune) int {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = XqkRuneListGetIndexByList(list, v)
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

// XqkRuneListGetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
func XqkRuneListGetIndexAndSubBySList(list []rune, subListArr ...[]rune) (int, []rune) {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = XqkRuneListGetIndexByList(list, v)
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

// XqkRuneListGetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func XqkRuneListGetElByIndex(list []rune, index int) rune {
	if XqkRuneListIsEmpty(list) || index < 0 || index >= len(list) {
		return 0
	}
	return list[index]
}

// XqkRuneListGetMax 获取切片中最大的byte，nil、空切片都会报错
func XqkRuneListGetMax(list []rune) (rune, error) {
	if len(list) == 0 {
		return 0, xqkVar.BaseErrListEmpty
	}
	max := list[0]
	for k := 1; k < len(list); k++ {
		if list[k] > max {
			max = list[k]
		}
	}
	return max, nil
}

// XqkRuneListGetMin 获取切片中最小的byte，nil、空切片都会报错
func XqkRuneListGetMin(list []rune) (rune, error) {
	if len(list) == 0 {
		return 0, xqkVar.BaseErrListEmpty
	}
	min := list[0]
	for k := 1; k < len(list); k++ {
		if list[k] < min {
			min = list[k]
		}
	}
	return min, nil
}
