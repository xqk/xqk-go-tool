package xqkAll

func XqkSIntToStr(iList ...int) string {
	return XqkSIntToStrBySep("", iList...)
}

func XqkSIntToStrBySep(sep string, iList ...int) string {
	return XqkIntListToStrBySep(iList, sep)
}
