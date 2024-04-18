package xqkVar

import (
	"errors"

	xqkLang "github.com/xqk/xqk-go-tool/xqk_lang"
)

var (
	BaseErrStrEmpty        = errors.New(xqkLang.GetErrStrEmptyLang())
	BaseErrListEmpty       = errors.New(xqkLang.GetErrListEmptyLang())
	BaseErrAddrNil         = errors.New(xqkLang.GetErrAddrNilLang())
	BaseErrIndexOutOfBound = errors.New(xqkLang.GetErrIndexOutOfBoundLang())
	BaseErrFileExists      = errors.New(xqkLang.GetErrFileExistsLang())
	BaseErrFileNotExists   = errors.New(xqkLang.GetErrFileNotExistsLang())
	BaseErrDirExists       = errors.New(xqkLang.GetErrDirExistsLang())
	BaseErrDirNotExists    = errors.New(xqkLang.GetErrDirNotExistsLang())
)
