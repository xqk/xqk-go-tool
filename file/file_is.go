package xqkFile

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// IsExists 判断路径指向的文件是否存在
func IsExists(path string) bool {
	return xqkAll.XqkFileIsExists(path)
}
