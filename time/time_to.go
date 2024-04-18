package xqkTime

import (
	"time"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

// ToStr1 格式化时间 => "2021-4-9 12:4:46"
// 其他格式参考 ToStrTransformMap
func ToStr1(t time.Time) string {
	return xqkAll.XqkTimeToStr1(t)
}

// ToStr10 格式化时间 => "2021-04-09 12:04:46.195482"
// 其他格式参考 ToStrTransformMap
func ToStr10(t time.Time) string {
	return xqkAll.XqkTimeToStr10(t)
}

// ToStr11 格式化时间 => "20214912446195482"
// 其他格式参考 ToStrTransformMap
func ToStr11(t time.Time) string {
	return xqkAll.XqkTimeToStr11(t)
}

// ToStr12 格式化时间 => "20210409120446195482"
// 其他格式参考 ToStrTransformMap
func ToStr12(t time.Time) string {
	return xqkAll.XqkTimeToStr12(t)
}

// ToStr13 格式化时间 => "2021年4月9日 12时4分46秒"
// 其他格式参考 ToStrTransformMap
func ToStr13(t time.Time) string {
	return xqkAll.XqkTimeToStr13(t)
}

// ToStr14 格式化时间 => "2021年04月09日 12时04分46秒"
// 其他格式参考 ToStrTransformMap
func ToStr14(t time.Time) string {
	return xqkAll.XqkTimeToStr14(t)
}

// ToStr15 格式化时间 => "2021年4月9日"
// 其他格式参考 ToStrTransformMap
func ToStr15(t time.Time) string {
	return xqkAll.XqkTimeToStr15(t)
}

// ToStr16 格式化时间 => "2021年04月09日"
// 其他格式参考 ToStrTransformMap
func ToStr16(t time.Time) string {
	return xqkAll.XqkTimeToStr16(t)
}

// ToStr17 格式化时间 => "April 9(st), 2021"
// 其他格式参考 ToStrTransformMap
func ToStr17(t time.Time) string {
	return xqkAll.XqkTimeToStr17(t)
}

// ToStr18 格式化时间 => "April 9, 2021"
// 其他格式参考 ToStrTransformMap
func ToStr18(t time.Time) string {
	return xqkAll.XqkTimeToStr18(t)
}

// ToStr19 格式化时间 => "April 9(st) 2021"
// 其他格式参考 ToStrTransformMap
func ToStr19(t time.Time) string {
	return xqkAll.XqkTimeToStr19(t)
}

// ToStr2 格式化时间 => "2021-04-09 12:04:46"
// 其他格式参考 ToStrTransformMap
func ToStr2(t time.Time) string {
	return xqkAll.XqkTimeToStr2(t)
}

// ToStr20 格式化时间 => "April 9 2021"
// 其他格式参考 ToStrTransformMap
func ToStr20(t time.Time) string {
	return xqkAll.XqkTimeToStr20(t)
}

// ToStr21 格式化时间 => "20210409"
// 其他格式参考 ToStrTransformMap
func ToStr21(t time.Time) string {
	return xqkAll.XqkTimeToStr21(t)
}

// ToStr22 格式化时间 => "120446"
// 其他格式参考 ToStrTransformMap
func ToStr22(t time.Time) string {
	return xqkAll.XqkTimeToStr22(t)
}

// ToStr3 格式化时间 => "2021-4-9"
// 其他格式参考 ToStrTransformMap
func ToStr3(t time.Time) string {
	return xqkAll.XqkTimeToStr3(t)
}

// ToStr4 格式化时间 => "2021-04-09"
// 其他格式参考 ToStrTransformMap
func ToStr4(t time.Time) string {
	return xqkAll.XqkTimeToStr4(t)
}

// ToStr5 格式化时间 => "12:4:46"
// 其他格式参考 ToStrTransformMap
func ToStr5(t time.Time) string {
	return xqkAll.XqkTimeToStr5(t)
}

// ToStr6 格式化时间 => "12:04:46"
// 其他格式参考 ToStrTransformMap
func ToStr6(t time.Time) string {
	return xqkAll.XqkTimeToStr6(t)
}

// ToStr7 格式化时间 => "20214912446"
// 其他格式参考 ToStrTransformMap
func ToStr7(t time.Time) string {
	return xqkAll.XqkTimeToStr7(t)
}

// ToStr8 格式化时间 => "20210409120446"
// 其他格式参考 ToStrTransformMap
func ToStr8(t time.Time) string {
	return xqkAll.XqkTimeToStr8(t)
}

// ToStr9 格式化时间 => "2021-4-9 12:4:46.195482"
// 其他格式参考 ToStrTransformMap
func ToStr9(t time.Time) string {
	return xqkAll.XqkTimeToStr9(t)
}

// ToStrTransformMap 所有ToStr函数 输出格式
func ToStrTransformMap() map[string]string {
	return xqkAll.XqkTimeToStrTransformMap()
}
