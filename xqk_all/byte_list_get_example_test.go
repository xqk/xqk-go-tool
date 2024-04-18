package xqkAll_test

import (
	"fmt"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func ExampleXqkByteListGetDeduplicate() {
	fmt.Println(xqkAll.XqkByteListGetDeduplicate([]byte{1, 2, 3, 1, 5, 4}))
	fmt.Println(xqkAll.XqkByteListGetDeduplicate([]byte{}))
	fmt.Println(xqkAll.XqkByteListGetDeduplicate(nil))
	// Output:
	// [1 2 3 5 4]
	// []
	// []
}

func ExampleXqkByteListGetDeleteByIndex() {
	fmt.Println(xqkAll.XqkByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, 0))
	fmt.Println(xqkAll.XqkByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, 2))
	fmt.Println(xqkAll.XqkByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, -11))
	fmt.Println(xqkAll.XqkByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, 999))
	fmt.Println(xqkAll.XqkByteListGetDeleteByIndex([]byte{}, 2))
	fmt.Println(xqkAll.XqkByteListGetDeleteByIndex(nil, 2))
	// Output:
	// [2 3 4]
	// [1 2 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleXqkByteListGetDeleteByRangeIndex() {
	fmt.Println(xqkAll.XqkByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 0))
	fmt.Println(xqkAll.XqkByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 1))
	fmt.Println(xqkAll.XqkByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, -2, 1))
	fmt.Println(xqkAll.XqkByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 2, 999))
	fmt.Println(xqkAll.XqkByteListGetDeleteByRangeIndex([]byte{}, 0, 1))
	fmt.Println(xqkAll.XqkByteListGetDeleteByRangeIndex(nil, 0, 2))
	// Output:
	// [2 3 4]
	// [3 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleXqkByteListGetFilter() {
	fmt.Println(xqkAll.XqkByteListGetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return true }))
	fmt.Println(xqkAll.XqkByteListGetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return false }))
	fmt.Println(xqkAll.XqkByteListGetFilter([]byte{}, func(byte) bool { return true }))
	fmt.Println(xqkAll.XqkByteListGetFilter([]byte{}, func(byte) bool { return false }))
	fmt.Println(xqkAll.XqkByteListGetFilter(nil, func(byte) bool { return true }))
	fmt.Println(xqkAll.XqkByteListGetFilter(nil, func(byte) bool { return false }))
	// Output:
	// [1 2 3 4]
	// []
	// []
	// []
	// []
	// []
}

func ExampleXqkByteListGetPop() {
	fmt.Println(xqkAll.XqkByteListGetPop([]byte{1, 2, 3, 4}))
	fmt.Println(xqkAll.XqkByteListGetPop([]byte{}))
	fmt.Println(xqkAll.XqkByteListGetPop(nil))
	// Output:
	// 4 [1 2 3] <nil>
	// 0 [] 切片不能为空
	// 0 [] 切片不能为空
}

func ExampleXqkByteListGetReverse() {
	fmt.Println(xqkAll.XqkByteListGetReverse([]byte{1, 2, 3, 4}))
	fmt.Println(xqkAll.XqkByteListGetReverse([]byte{}))
	fmt.Println(xqkAll.XqkByteListGetReverse(nil))
	// Output:
	// [4 3 2 1]
	// []
	// []
}

func ExampleXqkByteListGetMap() {
	fmt.Println(xqkAll.XqkByteListGetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i+1) + b }))
	fmt.Println(xqkAll.XqkByteListGetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(xqkAll.XqkByteListGetMap([]byte{}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(xqkAll.XqkByteListGetMap(nil, func(i int, b byte) byte { return byte(i) }))
	// Output:
	// [5 5 5 5]
	// [0 1 2 3]
	// []
	// []
}

func ExampleXqkByteListGetCopy() {
	fmt.Println(xqkAll.XqkByteListGetCopy([]byte{1, 2, 3, 4}))
	fmt.Println(xqkAll.XqkByteListGetCopy([]byte{}))
	fmt.Println(xqkAll.XqkByteListGetCopy(nil))
	// Output:
	// [1 2 3 4]
	// []
	// []
}

func ExampleXqkByteListGetIndexByEl() {
	fmt.Println(xqkAll.XqkByteListGetIndexByEl([]byte{1, 2, 3, 4}, 1))
	fmt.Println(xqkAll.XqkByteListGetIndexByEl([]byte{1, 2, 3, 4}, 9))
	fmt.Println(xqkAll.XqkByteListGetIndexByEl([]byte{}, 1))
	fmt.Println(xqkAll.XqkByteListGetIndexByEl(nil, 1))
	// Output:
	// 0
	// -1
	// -1
	// -1
}

func ExampleXqkByteListGetIndexByList() {
	fmt.Println(xqkAll.XqkByteListGetIndexByList([]byte{1, 2, 3, 4}, []byte{2, 3}))
	fmt.Println(xqkAll.XqkByteListGetIndexByList([]byte{1, 2, 3, 4}, []byte{}))
	fmt.Println(xqkAll.XqkByteListGetIndexByList([]byte{1, 2, 3, 4}, []byte{4, 3}))
	fmt.Println(xqkAll.XqkByteListGetIndexByList([]byte{1, 2, 3, 4}, nil))
	fmt.Println(xqkAll.XqkByteListGetIndexByList([]byte{}, []byte{2, 3}))
	fmt.Println(xqkAll.XqkByteListGetIndexByList([]byte{}, []byte{}))
	fmt.Println(xqkAll.XqkByteListGetIndexByList(nil, []byte{2, 3}))
	fmt.Println(xqkAll.XqkByteListGetIndexByList(nil, nil))
	// Output:
	// 1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
}
