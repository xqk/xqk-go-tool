package xqkBool_test

import (
	"fmt"

	xqkBool "github.com/xqk/xqk-go-tool/bool"
)

func ExampleToInt() {
	fmt.Println(xqkBool.ToInt(true))
	fmt.Println(xqkBool.ToInt(false))
	// Output:
	// 1
	// 0
}

func ExampleToStr() {
	fmt.Println(xqkBool.ToStr(true))
	fmt.Println(xqkBool.ToStr(false))
	// Output:
	// true
	// false
}
