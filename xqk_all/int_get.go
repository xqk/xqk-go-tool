package xqkAll

// XqkIntGetTrueInt 获取 true 的int值
func XqkIntGetTrueInt() int {
	return 1
}

// XqkIntGetFalseInt 获取 false 的int值
func XqkIntGetFalseInt() int {
	return 0
}

// XqkIntGetEnglishNumTail 获取英文的数字后缀
func XqkIntGetEnglishNumTail(i int) string {
	if i == 1 {
		return "st"
	}
	if i == 2 {
		return "nd"
	}
	if i == 3 {
		return "rd"
	}
	return "th"
}
