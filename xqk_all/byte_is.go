package xqkAll

// XqkByteIsLetter 判断是否是字母`a-zA-Z`，如果是返回`true`
func XqkByteIsLetter(b byte) bool {
	return XqkByteIsLowerLetter(b) || XqkByteIsUpperLetter(b)
}

// XqkByteIsLowerLetter 判断是否是小写字母`a-z`，如果是返回`true`
func XqkByteIsLowerLetter(b byte) bool {
	return 97 <= b && b <= 122
}

// XqkByteIsUpperLetter 判断是否是大写字母`A-Z`，如果是返回`true`
func XqkByteIsUpperLetter(b byte) bool {
	return 65 <= b && b <= 90
}

// XqkByteIsNotLetter 判断是否不是字母`a-zA-Z`，如果不是字母返回`true`
func XqkByteIsNotLetter(b byte) bool {
	return !XqkByteIsLetter(b)
}
