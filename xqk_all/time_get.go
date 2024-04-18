package xqkAll

import (
	"time"
)

// XqkTimeGetNowStrTransformMap 所有GetNowStr函数 输出格式
func XqkTimeGetNowStrTransformMap() map[string]string {
	return map[string]string{
		"GetNowStr1":  "2021-4-9 12:4:46",
		"GetNowStr2":  "2021-04-09 12:04:46",
		"GetNowStr3":  "2021-4-9",
		"GetNowStr4":  "2021-04-09",
		"GetNowStr5":  "12:4:46",
		"GetNowStr6":  "12:04:46",
		"GetNowStr7":  "20214912446",
		"GetNowStr8":  "20210409120446",
		"GetNowStr9":  "2021-4-9 12:4:46.195482",
		"GetNowStr10": "2021-04-09 12:04:46.195482",
		"GetNowStr11": "20214912446195482",
		"GetNowStr12": "20210409120446195482",
		"GetNowStr13": "2021年4月9日 12时4分46秒",
		"GetNowStr14": "2021年04月09日 12时04分46秒",
		"GetNowStr15": "2021年4月9日",
		"GetNowStr16": "2021年04月09日",
		"GetNowStr17": "April 9(st), 2021",
		"GetNowStr18": "April 9, 2021",
		"GetNowStr19": "April 9(st) 2021",
		"GetNowStr20": "April 9 2021",
		"GetNowStr21": "20210409",
		"GetNowStr22": "120446",
	}
}

// XqkTimeGetMonthEnglish 获取月份的英文单词
func XqkTimeGetMonthEnglish(m time.Month) string {
	switch m {
	case time.January:
		return "January"
	case time.February:
		return "February"
	case time.March:
		return "March"
	case time.April:
		return "April"
	case time.May:
		return "May"
	case time.June:
		return "June"
	case time.July:
		return "July"
	case time.August:
		return "August"
	case time.September:
		return "September"
	case time.October:
		return "October"
	case time.November:
		return "November"
	case time.December:
		return "December"
	default:
		return "January"
	}
}

// XqkTimeGetNowStr1 获取当前时间字符串 "2021-4-9 12:4:46"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr1() string {
	now := time.Now()
	return XqkTimeToStr1(now)
}

// XqkTimeGetNowStr2 获取当前时间字符串 "2021-04-09 12:04:46"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr2() string {
	now := time.Now()
	return XqkTimeToStr2(now)
}

// XqkTimeGetNowStr3 获取当前时间字符串 "2021-4-9"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr3() string {
	now := time.Now()
	return XqkTimeToStr3(now)
}

// XqkTimeGetNowStr4 获取当前时间字符串 "2021-04-09"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr4() string {
	now := time.Now()
	return XqkTimeToStr4(now)
}

// XqkTimeGetNowStr5 获取当前时间字符串 "12:4:46"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr5() string {
	now := time.Now()
	return XqkTimeToStr5(now)
}

// XqkTimeGetNowStr6 获取当前时间字符串 "12:04:46"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr6() string {
	now := time.Now()
	return XqkTimeToStr6(now)
}

// XqkTimeGetNowStr7 获取当前时间字符串 "20214912446"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr7() string {
	now := time.Now()
	return XqkTimeToStr7(now)
}

// XqkTimeGetNowStr8 获取当前时间字符串 "20210409120446"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr8() string {
	now := time.Now()
	return XqkTimeToStr8(now)
}

// XqkTimeGetNowStr9 获取当前时间字符串 "2021-4-9 12:4:46.195482"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr9() string {
	now := time.Now()
	return XqkTimeToStr9(now)
}

// XqkTimeGetNowStr10 获取当前时间字符串 "2021-04-09 12:04:46.195482"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr10() string {
	now := time.Now()
	return XqkTimeToStr10(now)
}

// XqkTimeGetNowStr11 获取当前时间字符串 "20214912446195482"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr11() string {
	now := time.Now()
	return XqkTimeToStr11(now)
}

// XqkTimeGetNowStr12 获取当前时间字符串 "20210409120446195482"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr12() string {
	now := time.Now()
	return XqkTimeToStr12(now)
}

// XqkTimeGetNowStr13 获取当前时间字符串 "2021年4月9日 12时4分46秒"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr13() string {
	now := time.Now()
	return XqkTimeToStr13(now)
}

// XqkTimeGetNowStr14 获取当前时间字符串 "2021年04月09日 12时04分46秒"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr14() string {
	now := time.Now()
	return XqkTimeToStr14(now)
}

// XqkTimeGetNowStr15 获取当前时间字符串 "2021年4月9日"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr15() string {
	now := time.Now()
	return XqkTimeToStr15(now)
}

// XqkTimeGetNowStr16 获取当前时间字符串 "2021年04月09日"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr16() string {
	now := time.Now()
	return XqkTimeToStr16(now)
}

// XqkTimeGetNowStr17 获取当前时间字符串 "April 9(st), 2021"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr17() string {
	now := time.Now()
	return XqkTimeToStr17(now)
}

// XqkTimeGetNowStr18 获取当前时间字符串 "April 9, 2021"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr18() string {
	now := time.Now()
	return XqkTimeToStr18(now)
}

// XqkTimeGetNowStr19 获取当前时间字符串 "April 9(st) 2021"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr19() string {
	now := time.Now()
	return XqkTimeToStr19(now)
}

// XqkTimeGetNowStr20 获取当前时间字符串 "April 9 2021"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr20() string {
	now := time.Now()
	return XqkTimeToStr20(now)
}

// XqkTimeGetNowStr21 获取当前时间字符串 "20210409"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr21() string {
	now := time.Now()
	return XqkTimeToStr21(now)
}

// XqkTimeGetNowStr22 获取当前时间字符串 "120446"
// 其他格式参考 GetNowStrTransformMap
func XqkTimeGetNowStr22() string {
	now := time.Now()
	return XqkTimeToStr22(now)
}
