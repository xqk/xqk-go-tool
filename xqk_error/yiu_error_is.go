package xqkBaseErr

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

func IsAddrNil(err error) bool {
	return xqkAll.XqkBaseErrIsAddrNil(err)
}

func IsDirExists(err error) bool {
	return xqkAll.XqkBaseErrIsDirExists(err)
}

func IsDirNotExists(err error) bool {
	return xqkAll.XqkBaseErrIsDirNotExists(err)
}

func IsFileExists(err error) bool {
	return xqkAll.XqkBaseErrIsFileExists(err)
}

func IsFileNotExists(err error) bool {
	return xqkAll.XqkBaseErrIsFileNotExists(err)
}

func IsIndexOutOfBound(err error) bool {
	return xqkAll.XqkBaseErrIsIndexOutOfBound(err)
}

func IsListEmpty(err error) bool {
	return xqkAll.XqkBaseErrIsListEmpty(err)
}

func IsStrEmpty(err error) bool {
	return xqkAll.XqkBaseErrIsStrEmpty(err)
}
