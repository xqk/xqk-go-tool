package xqkAll

import (
	"os"
)

// XqkFileIsExists 判断路径指向的文件是否存在
func XqkFileIsExists(path string) bool {
	osInfo, err := os.Stat(path)
	if err == nil {
		// 路径指向有效
		return !osInfo.IsDir()
	}
	// return !os.IsNotExist(err)
	// 此时就算文件存在，但是还是不能确定是文件
	return false
}
