package xqkErr

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// IsExitStatus1 判断错误信息是否是 "exit status 1"
func IsExitStatus1(err error) bool {
	return xqkAll.XqkErrIsExitStatus1(err)
}
