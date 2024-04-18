package xqkAll

import xqkVar "github.com/xqk/xqk-go-tool/xqk_var"

// XqkStrListOpDeduplicate 去重，按顺序保留
func XqkStrListOpDeduplicate(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetDeduplicate(*list)
}

// XqkStrListOpDeleteByIndex 根据索引删除，超出索引不处理
func XqkStrListOpDeleteByIndex(list *[]string, delIndex int) {
	if list == nil {
		return
	}
	*list = XqkStrListGetDeleteByIndex(*list, delIndex)
}

// XqkStrListOpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func XqkStrListOpDeleteByRangeIndex(list *[]string, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = XqkStrListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// XqkStrListOpFilter 过滤切片，保留返回ture的
func XqkStrListOpFilter(list *[]string, keep func(x string) bool) {
	if list == nil {
		return
	}
	*list = XqkStrListGetFilter(*list, keep)
}

// XqkStrListOpPop 切片元素出栈，nil、空切片都会报错
func XqkStrListOpPop(list *[]string) (string, error) {
	if list == nil {
		return "", xqkVar.BaseErrAddrNil
	}
	pop, tempList, err := XqkStrListGetPop(*list)
	if err != nil {
		return "", err
	}
	*list = tempList
	return pop, nil
}

// XqkStrListOpReverse 切片元素反转
func XqkStrListOpReverse(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetReverse(*list)
}

// XqkStrListOpShuffle 切片元素乱序排列
func XqkStrListOpShuffle(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetShuffle(*list)
}

// XqkStrListOpMap 遍历计算切片，修改原切片
//
// opFunc =
// func(i int, s string) string {
// 	   return s + "(" + XqkIntUtil.ToStr(i) + ")"
// }
//
// ["Hello", "Fidel", "Xqk"] > opFunc > ["Hello(0)" "Fidel(1)" "Xqk(2)"]
//
// [] > opFunc > []
//
// nil > opFunc > nil
func XqkStrListOpMap(list *[]string, opFunc func(int, string) string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetMap(*list, opFunc)
}

// XqkStrListOpDeleteEmptyEl 去除所有空串
//
// ["", "Hello", "", "Xqk"] >> ["Hello", "Xqk"]
//
// [" ", "Hello", "", "Xqk"] >> [" ", "Hello", "Xqk"]
//
// [] >> []
//
// nil >> nil
func XqkStrListOpDeleteEmptyEl(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetDeleteEmptyEl(*list)
}

// XqkStrListOpDeleteBlankEl 去除所有空串和空格字符串
//
// ["", "Hello", "", "Xqk"] >> ["Hello", "Xqk"]
//
// [" ", "Hello", "", "Xqk"] >> ["Hello", "Xqk"]
//
// [] >> []
//
// nil >> nil
func XqkStrListOpDeleteBlankEl(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetDeleteBlankEl(*list)
}

// XqkStrListOpAscUnicode 按照Unicode逐个升序排列
func XqkStrListOpAscUnicode(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetAscUnicode(*list)
}

// XqkStrListOpDescUnicode 按照Unicode逐个将序排列
func XqkStrListOpDescUnicode(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkStrListGetDescUnicode(*list)
}
