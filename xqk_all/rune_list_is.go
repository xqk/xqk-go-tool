package xqkAll

import "unicode"

// XqkRuneListIsContainsEl 判断切片是否包含某byte
func XqkRuneListIsContainsEl(list []rune, x rune) bool {
	if len(list) == 0 {
		return false
	}
	for k := range list {
		if list[k] == x {
			return true
		}
	}
	return false
}

// XqkRuneListIsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
func XqkRuneListIsContainsElList(list, subList []rune) bool {
	if XqkRuneListGetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// XqkRuneListIsNil 判断切片是否为nil
func XqkRuneListIsNil(list []rune) bool {
	return list == nil
}

// XqkRuneListIsEmpty 判断切片长度是否等于0
func XqkRuneListIsEmpty(list []rune) bool {
	return len(list) == 0
}

// XqkRuneListIsEqual 判断两个数组是否相等
func XqkRuneListIsEqual(listA, listB []rune) bool {
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

// XqkRuneListIsGtByUnicode 按unicode逐个比较，listA是否大于listB
func XqkRuneListIsGtByUnicode(listA, listB []rune) bool {
	for i := range listA {
		itemB := XqkRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		return listA[i] > itemB
	}
	return true
}

// XqkRuneListIsGeByUnicode 按unicode逐个比较，listA是否大于等于listB
func XqkRuneListIsGeByUnicode(listA, listB []rune) bool {
	if XqkRuneListIsEqual(listA, listB) {
		return true
	}
	return XqkRuneListIsGtByUnicode(listA, listB)
}

// XqkRuneListIsLtByUnicode 按unicode逐个比较，listA是否小于listB
func XqkRuneListIsLtByUnicode(listA, listB []rune) bool {
	for i := range listA {
		itemB := XqkRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		return listA[i] < itemB
	}
	return true
}

// XqkRuneListIsLeByUnicode 按unicode逐个比较，listA是否小于等于listB
func XqkRuneListIsLeByUnicode(listA, listB []rune) bool {
	if XqkRuneListIsEqual(listA, listB) {
		return true
	}
	return XqkRuneListIsLtByUnicode(listA, listB)
}

// XqkRuneListIsGtByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否大于listB
//
// 其实就是将大小写颠倒一下
func XqkRuneListIsGtByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	for i := range listA {
		itemB := XqkRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		if unicode.IsLetter(listA[i]) && unicode.IsLetter(itemB) {
			// 如果都是字母
			if (unicode.IsUpper(listA[i]) && unicode.IsUpper(itemB)) ||
				(unicode.IsLower(listA[i]) && unicode.IsLower(itemB)) {
				// 同为大小写
				return listA[i] > itemB
			} else {
				// 不同为大小写，只要是小写就是大于
				return unicode.IsUpper(listA[i])
			}
		} else {
			// 不都是字母，那么直接使用unicode排序
			return listA[i] > itemB
		}
	}
	return false
}

// XqkRuneListIsGeByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否大于等于listB
//
// 其实就是将大小写颠倒一下
func XqkRuneListIsGeByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	if XqkRuneListIsEqual(listA, listB) {
		return true
	}
	return XqkRuneListIsGtByUnicodeAndLowerBeforeUpper(listA, listB)
}

// XqkRuneListIsLtByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否小于listB
//
// 其实就是将大小写颠倒一下
func XqkRuneListIsLtByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	for i := range listA {
		itemB := XqkRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		if unicode.IsLetter(listA[i]) && unicode.IsLetter(itemB) {
			// 如果都是字母
			if (unicode.IsUpper(listA[i]) && unicode.IsUpper(itemB)) ||
				(unicode.IsLower(listA[i]) && unicode.IsLower(itemB)) {
				// 同为大小写
				return listA[i] < itemB
			} else {
				// 不同为大小写，只要是大写就是小于
				return unicode.IsLower(listA[i])
			}
		} else {
			// 不都是字母，那么直接使用unicode排序
			return listA[i] < itemB
		}
	}
	return false
}

// XqkRuneListIsLeByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否小于等于listB
//
// 其实就是将大小写颠倒一下
func XqkRuneListIsLeByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	if XqkRuneListIsEqual(listA, listB) {
		return true
	}
	return XqkRuneListIsLtByUnicodeAndLowerBeforeUpper(listA, listB)
}
