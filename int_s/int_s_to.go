package xqkSInt

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

func ToStr(iList ...int) string {
	return xqkAll.XqkSIntToStr(iList...)
}

func ToStrBySep(sep string, iList ...int) string {
	return xqkAll.XqkSIntToStrBySep(sep, iList...)
}
