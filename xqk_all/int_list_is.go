package xqkAll

import "github.com/psampaz/slice"

// XqkIntListIsContainsEl 判断切片是否包含某byte
func XqkIntListIsContainsEl(list []int, b int) bool {
	return slice.ContainsInt(list, b)
}

// XqkIntListIsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
//
// [1, 2, 3, 4] > [2, 3] > true
//
// [1, 2, 3, 4] > [5, 3] > false
//
// [] > [2, 3] > false
//
// nil > [2, 3] > false
//
// [1, 2] > [] > false
//
// [1, 2] > nil > false
//
// [] > [] > false
//
// nil > nil > false
func XqkIntListIsContainsElList(list, subList []int) bool {
	if XqkIntListGetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// XqkIntListIsNil 判断切片是否为nil
func XqkIntListIsNil(list []int) bool {
	return list == nil
}

// XqkIntListIsEmpty 判断切片长度是否等于0
func XqkIntListIsEmpty(list []int) bool {
	return len(list) == 0
}

// XqkIntListIsEqual 判断两个数组是否相等
func XqkIntListIsEqual(listA, listB []int) bool {
	if (listA == nil) != (listB == nil) {
		return false
	}
	if len(listA) != len(listB) {
		return false
	}
	for i := range listA {
		if listA[i] != listB[i] {
			return false
		}
	}
	return true
}
