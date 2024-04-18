package xqkSErr

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// ToError 将所有错误信息合并后生成一个新的error
func ToError(errList ...error) error {
	return xqkAll.XqkSErrToError(errList...)
}

// ToErrorBySep 将所有错误信息按分隔符合并后生成一个新的error
func ToErrorBySep(sep string, errList ...error) error {
	return xqkAll.XqkSErrToErrorBySep(sep, errList...)
}

// ToStr 将所有错误的信息合并输出
func ToStr(errList ...error) string {
	return xqkAll.XqkSErrToStr(errList...)
}

// ToStrBySep 将所有错误的信息按分隔符合并输出
func ToStrBySep(sep string, errList ...error) string {
	return xqkAll.XqkSErrToStrBySep(sep, errList...)
}
