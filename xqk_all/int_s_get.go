package xqkAll

// XqkSIntGetMax 获取参数中最大的值
func XqkSIntGetMax(iList ...int) (int, error) {
	return XqkIntListGetMax(iList)
}

// XqkSIntGetMin 获取参数中最小的值
func XqkSIntGetMin(iList ...int) (int, error) {
	return XqkIntListGetMin(iList)
}

// XqkSIntGetSum 参数求和
func XqkSIntGetSum(iList ...int) int {
	return XqkIntListGetSum(iList)
}
