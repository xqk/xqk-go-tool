package xqkAll

// XqkSByteListGetMerge 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func XqkSByteListGetMerge(bListArr ...[]byte) []byte {
	outLength := 0
	for _, v := range bListArr {
		outLength += len(v)
	}
	if outLength == 0 {
		return nil
	}
	outList := make([]byte, outLength, outLength)
	mergeIndex := 0
	for _, v := range bListArr {
		copy(outList[mergeIndex:mergeIndex+len(v)], v)
		mergeIndex += len(v)
	}
	return outList
}

// XqkSByteListGetMaxLengthEl 获取长度最大的第一个切片
func XqkSByteListGetMaxLengthEl(list ...[]byte) []byte {
	if len(list) == 0 {
		return nil
	}
	resultIndex := 0
	maxLength := len(list[0])
	for i := range list {
		if maxLength < len(list[i]) {
			resultIndex = i
			maxLength = len(list[i])
		}
	}
	return list[resultIndex]
}

// XqkSByteListGetMinLengthEl 获取长度最小的第一个切片
func XqkSByteListGetMinLengthEl(list ...[]byte) []byte {
	if len(list) == 0 {
		return nil
	}
	resultIndex := 0
	maxLength := len(list[0])
	for i := range list {
		if maxLength > len(list[i]) {
			resultIndex = i
			maxLength = len(list[i])
		}
	}
	return list[resultIndex]
}
