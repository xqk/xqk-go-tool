package xqkByteList_test

import (
	"fmt"

	xqkByteList "github.com/xqk/xqk-go-tool/byte_list"
)

func ExampleGetDeduplicate() {
	fmt.Println(xqkByteList.GetDeduplicate([]byte{1, 2, 3, 1, 5, 4}))
	fmt.Println(xqkByteList.GetDeduplicate([]byte{}))
	fmt.Println(xqkByteList.GetDeduplicate(nil))
	// Output:
	// [1 2 3 5 4]
	// []
	// []
}

func ExampleGetDeleteByIndex() {
	fmt.Println(xqkByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, 0))
	fmt.Println(xqkByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, 2))
	fmt.Println(xqkByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, -11))
	fmt.Println(xqkByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, 999))
	fmt.Println(xqkByteList.GetDeleteByIndex([]byte{}, 2))
	fmt.Println(xqkByteList.GetDeleteByIndex(nil, 2))
	// Output:
	// [2 3 4]
	// [1 2 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleGetDeleteByRangeIndex() {
	fmt.Println(xqkByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 0))
	fmt.Println(xqkByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 1))
	fmt.Println(xqkByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, -2, 1))
	fmt.Println(xqkByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 2, 999))
	fmt.Println(xqkByteList.GetDeleteByRangeIndex([]byte{}, 0, 1))
	fmt.Println(xqkByteList.GetDeleteByRangeIndex(nil, 0, 2))
	// Output:
	// [2 3 4]
	// [3 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleGetFilter() {
	fmt.Println(xqkByteList.GetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return true }))
	fmt.Println(xqkByteList.GetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return false }))
	fmt.Println(xqkByteList.GetFilter([]byte{}, func(byte) bool { return true }))
	fmt.Println(xqkByteList.GetFilter([]byte{}, func(byte) bool { return false }))
	fmt.Println(xqkByteList.GetFilter(nil, func(byte) bool { return true }))
	fmt.Println(xqkByteList.GetFilter(nil, func(byte) bool { return false }))
	// Output:
	// [1 2 3 4]
	// []
	// []
	// []
	// []
	// []
}

func ExampleGetPop() {
	fmt.Println(xqkByteList.GetPop([]byte{1, 2, 3, 4}))
	fmt.Println(xqkByteList.GetPop([]byte{}))
	fmt.Println(xqkByteList.GetPop(nil))
	// Output:
	// 4 [1 2 3] <nil>
	// 0 [] 切片不能为空
	// 0 [] 切片不能为空
}

func ExampleGetReverse() {
	fmt.Println(xqkByteList.GetReverse([]byte{1, 2, 3, 4}))
	fmt.Println(xqkByteList.GetReverse([]byte{}))
	fmt.Println(xqkByteList.GetReverse(nil))
	// Output:
	// [4 3 2 1]
	// []
	// []
}

func ExampleGetMap() {
	fmt.Println(xqkByteList.GetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i+1) + b }))
	fmt.Println(xqkByteList.GetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(xqkByteList.GetMap([]byte{}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(xqkByteList.GetMap(nil, func(i int, b byte) byte { return byte(i) }))
	// Output:
	// [5 5 5 5]
	// [0 1 2 3]
	// []
	// []
}

func ExampleGetCopy() {
	fmt.Println(xqkByteList.GetCopy([]byte{1, 2, 3, 4}))
	fmt.Println(xqkByteList.GetCopy([]byte{}))
	fmt.Println(xqkByteList.GetCopy(nil))
	// Output:
	// [1 2 3 4]
	// []
	// []
}

func ExampleGetIndexByEl() {
	fmt.Println(xqkByteList.GetIndexByEl([]byte{1, 2, 3, 4}, 1))
	fmt.Println(xqkByteList.GetIndexByEl([]byte{1, 2, 3, 4}, 9))
	fmt.Println(xqkByteList.GetIndexByEl([]byte{}, 1))
	fmt.Println(xqkByteList.GetIndexByEl(nil, 1))
	// Output:
	// 0
	// -1
	// -1
	// -1
}

func ExampleGetIndexByList() {
	fmt.Println(xqkByteList.GetIndexByList([]byte{1, 2, 3, 4}, []byte{2, 3}))
	fmt.Println(xqkByteList.GetIndexByList([]byte{1, 2, 3, 4}, []byte{}))
	fmt.Println(xqkByteList.GetIndexByList([]byte{1, 2, 3, 4}, []byte{4, 3}))
	fmt.Println(xqkByteList.GetIndexByList([]byte{1, 2, 3, 4}, nil))
	fmt.Println(xqkByteList.GetIndexByList([]byte{}, []byte{2, 3}))
	fmt.Println(xqkByteList.GetIndexByList([]byte{}, []byte{}))
	fmt.Println(xqkByteList.GetIndexByList(nil, []byte{2, 3}))
	fmt.Println(xqkByteList.GetIndexByList(nil, nil))
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
