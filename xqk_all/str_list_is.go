package xqkAll

import (
	"strings"

	"github.com/psampaz/slice"
)

// XqkStrListIsContainsEl 判断切片是否包含某字符串
func XqkStrListIsContainsEl(list []string, str string) bool {
	return slice.ContainsString(list, str)
}

// XqkStrListIsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
//
// ["a", "b", "c", "d"] > ["b", "c"] > true
//
// ["a", "b", "c", "d"] > ["y", "c"] > false
//
// [] > ["y", "c"] > false
//
// nil > ["y", "c"] > false
//
// ["a", "b"] > [] > false
//
// ["a", "b"] > nil > false
//
// [] > [] > false
//
// nil > nil > false
func XqkStrListIsContainsElList(list, subList []string) bool {
	if XqkStrListGetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// XqkStrListIsNil 判断切片是否为nil
func XqkStrListIsNil(list []string) bool {
	return list == nil
}

// XqkStrListIsEmpty 判断切片长度是否等于0
func XqkStrListIsEmpty(list []string) bool {
	return len(list) == 0
}

// XqkStrListIsEqual 判断两个数组是否相等
func XqkStrListIsEqual(listA, listB []string) bool {
	return xqkStrListIsEqual(listA, listB, func(strA, strB string) bool {
		return strA == strB
	})
}

// XqkStrListIsEqualFold 按utf-8编码判断是否相等，不区分大小写
func XqkStrListIsEqualFold(listA, listB []string) bool {
	return xqkStrListIsEqual(listA, listB, func(strA, strB string) bool {
		return strings.EqualFold(strA, strB)
	})
}

func xqkStrListIsEqual(listA, listB []string, equalFunc func(strA, strB string) bool) bool {
	if (listA == nil) != (listB == nil) {
		return false
	}
	if len(listA) != len(listB) {
		return false
	}
	for i := range listA {
		if !equalFunc(listA[i], listB[i]) {
			return false
		}
	}
	return true
}
