package xqkAll_test

import (
	"fmt"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func ExampleXqkBoolToInt() {
	fmt.Println(xqkAll.XqkBoolToInt(true))
	fmt.Println(xqkAll.XqkBoolToInt(false))
	// Output:
	// 1
	// 0
}

func ExampleXqkBoolToStr() {
	fmt.Println(xqkAll.XqkBoolToStr(true))
	fmt.Println(xqkAll.XqkBoolToStr(false))
	// Output:
	// true
	// false
}
