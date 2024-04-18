package xqkAll_test

import (
	"fmt"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func ExampleXqkByteIsLetter() {
	fmt.Println(xqkAll.XqkByteIsLetter('a'))
	fmt.Println(xqkAll.XqkByteIsLetter('B'))
	fmt.Println(xqkAll.XqkByteIsLetter('*'))
	// Output:
	// true
	// true
	// false
}

func ExampleXqkByteIsLowerLetter() {
	fmt.Println(xqkAll.XqkByteIsLowerLetter('a'))
	fmt.Println(xqkAll.XqkByteIsLowerLetter('B'))
	fmt.Println(xqkAll.XqkByteIsLowerLetter('*'))
	// Output:
	// true
	// false
	// false
}

func ExampleXqkByteIsUpperLetter() {
	fmt.Println(xqkAll.XqkByteIsUpperLetter('a'))
	fmt.Println(xqkAll.XqkByteIsUpperLetter('B'))
	fmt.Println(xqkAll.XqkByteIsUpperLetter('*'))
	// Output:
	// false
	// true
	// false
}

func ExampleXqkByteIsNotLetter() {
	fmt.Println(xqkAll.XqkByteIsNotLetter('a'))
	fmt.Println(xqkAll.XqkByteIsNotLetter('B'))
	fmt.Println(xqkAll.XqkByteIsNotLetter('*'))
	// Output:
	// false
	// false
	// true
}
