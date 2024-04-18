package xqkAll

import (
	"errors"

	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

func XqkBaseErrIsStrEmpty(err error) bool {
	return errors.Is(err, xqkVar.BaseErrStrEmpty)
}
func XqkBaseErrIsListEmpty(err error) bool {
	return errors.Is(err, xqkVar.BaseErrListEmpty)
}
func XqkBaseErrIsAddrNil(err error) bool {
	return errors.Is(err, xqkVar.BaseErrAddrNil)
}
func XqkBaseErrIsIndexOutOfBound(err error) bool {
	return errors.Is(err, xqkVar.BaseErrIndexOutOfBound)
}
func XqkBaseErrIsFileExists(err error) bool {
	return errors.Is(err, xqkVar.BaseErrFileExists)
}
func XqkBaseErrIsFileNotExists(err error) bool {
	return errors.Is(err, xqkVar.BaseErrFileNotExists)
}
func XqkBaseErrIsDirExists(err error) bool {
	return errors.Is(err, xqkVar.BaseErrDirExists)
}
func XqkBaseErrIsDirNotExists(err error) bool {
	return errors.Is(err, xqkVar.BaseErrDirNotExists)
}
