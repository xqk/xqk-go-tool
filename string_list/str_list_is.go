package xqkStrList

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// IsContainsEl 判断切片是否包含某字符串
func IsContainsEl(list []string, str string) bool {
	return xqkAll.XqkStrListIsContainsEl(list, str)
}

// IsContainsElList 判断切片是否包含子切片
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
func IsContainsElList(list, subList []string) bool {
	return xqkAll.XqkStrListIsContainsElList(list, subList)
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []string) bool {
	return xqkAll.XqkStrListIsEmpty(list)
}

// IsEqual 判断两个数组是否相等
func IsEqual(listA, listB []string) bool {
	return xqkAll.XqkStrListIsEqual(listA, listB)
}

// IsEqualFold 按utf-8编码判断是否相等，不区分大小写
func IsEqualFold(listA, listB []string) bool {
	return xqkAll.XqkStrListIsEqualFold(listA, listB)
}

// IsNil 判断切片是否为nil
func IsNil(list []string) bool {
	return xqkAll.XqkStrListIsNil(list)
}
