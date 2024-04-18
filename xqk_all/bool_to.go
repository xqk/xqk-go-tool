package xqkAll

// XqkBoolToInt bool转换成Int
func XqkBoolToInt(b bool) int {
	if b {
		return XqkIntGetTrueInt()
	} else {
		return XqkIntGetFalseInt()
	}
}

// XqkBoolToStr bool转换成string
func XqkBoolToStr(b bool) string {
	if b {
		return XqkStrGetTureStr()
	} else {
		return XqkStrGetFalseStr()
	}
}
