package xqkByte_test

import (
	"fmt"

	xqkByte "github.com/xqk/xqk-go-tool/byte"
)

func ExampleIsLetter() {
	fmt.Println(xqkByte.IsLetter('a'))
	fmt.Println(xqkByte.IsLetter('B'))
	fmt.Println(xqkByte.IsLetter('*'))
	// Output:
	// true
	// true
	// false
}

func ExampleIsLowerLetter() {
	fmt.Println(xqkByte.IsLowerLetter('a'))
	fmt.Println(xqkByte.IsLowerLetter('B'))
	fmt.Println(xqkByte.IsLowerLetter('*'))
	// Output:
	// true
	// false
	// false
}

func ExampleIsUpperLetter() {
	fmt.Println(xqkByte.IsUpperLetter('a'))
	fmt.Println(xqkByte.IsUpperLetter('B'))
	fmt.Println(xqkByte.IsUpperLetter('*'))
	// Output:
	// false
	// true
	// false
}

func ExampleIsNotLetter() {
	fmt.Println(xqkByte.IsNotLetter('a'))
	fmt.Println(xqkByte.IsNotLetter('B'))
	fmt.Println(xqkByte.IsNotLetter('*'))
	// Output:
	// false
	// false
	// true
}
