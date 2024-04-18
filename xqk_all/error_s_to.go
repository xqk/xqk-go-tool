package xqkAll

import (
	"errors"
)

// XqkSErrToStr 将所有错误的信息合并输出
func XqkSErrToStr(errList ...error) string {
	return XqkSErrToStrBySep("", errList...)
}

// XqkSErrToStrBySep 将所有错误的信息按分隔符合并输出
func XqkSErrToStrBySep(sep string, errList ...error) string {
	errStrList := make([]string, 0)
	for _, v := range errList {
		if v == nil {
			continue
		}
		errStrList = append(errStrList, v.Error())
	}
	if len(errStrList) > 0 {
		return XqkStrListToStrBySep(errStrList, sep)
	}
	return ""
}

// XqkSErrToError 将所有错误信息合并后生成一个新的error
func XqkSErrToError(errList ...error) error {
	return XqkSErrToErrorBySep("", errList...)
}

// XqkSErrToErrorBySep 将所有错误信息按分隔符合并后生成一个新的error
func XqkSErrToErrorBySep(sep string, errList ...error) error {
	errStr := XqkSErrToStrBySep(sep, errList...)
	if errStr == "" {
		return nil
	}
	return errors.New(errStr)
}
