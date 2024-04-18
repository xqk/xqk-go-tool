package xqkAll

import (
	"fmt"
	"time"
)

// XqkTimeToStrTransformMap 所有ToStr函数 输出格式
func XqkTimeToStrTransformMap() map[string]string {
	return map[string]string{
		"ToStr1":  "2021-4-9 12:4:46",
		"ToStr2":  "2021-04-09 12:04:46",
		"ToStr3":  "2021-4-9",
		"ToStr4":  "2021-04-09",
		"ToStr5":  "12:4:46",
		"ToStr6":  "12:04:46",
		"ToStr7":  "20214912446",
		"ToStr8":  "20210409120446",
		"ToStr9":  "2021-4-9 12:4:46.195482",
		"ToStr10": "2021-04-09 12:04:46.195482",
		"ToStr11": "20214912446195482",
		"ToStr12": "20210409120446195482",
		"ToStr13": "2021年4月9日 12时4分46秒",
		"ToStr14": "2021年04月09日 12时04分46秒",
		"ToStr15": "2021年4月9日",
		"ToStr16": "2021年04月09日",
		"ToStr17": "April 9(st), 2021",
		"ToStr18": "April 9, 2021",
		"ToStr19": "April 9(st) 2021",
		"ToStr20": "April 9 2021",
		"ToStr21": "20210409",
		"ToStr22": "120446",
	}
}

// XqkTimeToStr1 格式化时间 => "2021-4-9 12:4:46"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr1(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d %d:%d:%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// XqkTimeToStr2 格式化时间 => "2021-04-09 12:04:46"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr2(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// XqkTimeToStr3 格式化时间 => "2021-4-9"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr3(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d",
		t.Year(),
		t.Month(),
		t.Day())
	return nowStr
}

// XqkTimeToStr4 格式化时间 => "2021-04-09"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr4(t time.Time) string {
	return t.Format("2006-01-02")
}

// XqkTimeToStr5 格式化时间 => "12:4:46"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr5(t time.Time) string {
	nowStr := fmt.Sprintf("%d:%d:%d",
		t.Hour(),
		t.Minute(),
		t.Second())
	t.Nanosecond()
	return nowStr
}

// XqkTimeToStr6 格式化时间 => "12:04:46"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr6(t time.Time) string {
	return t.Format("15:04:05")
}

// XqkTimeToStr7 格式化时间 => "20214912446"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr7(t time.Time) string {
	nowStr := fmt.Sprintf("%d%d%d%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// XqkTimeToStr8 格式化时间 => "20210409120446"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr8(t time.Time) string {
	return t.Format("20060102150405")
}

// XqkTimeToStr9 格式化时间 => "2021-4-9 12:4:46.195482"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr9(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d %d:%d:%d.%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond())
	return nowStr
}

// XqkTimeToStr10 格式化时间 => "2021-04-09 12:04:46.195482"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr10(t time.Time) string {
	nowStr := fmt.Sprintf(t.Format("2006-01-02 15:04:05")+".%d",
		t.Nanosecond())
	return nowStr
}

// XqkTimeToStr11 格式化时间 => "20214912446195482"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr11(t time.Time) string {
	nowStr := fmt.Sprintf("%d%d%d%d%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond())
	return nowStr
}

// XqkTimeToStr12 格式化时间 => "20210409120446195482"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr12(t time.Time) string {
	nowStr := fmt.Sprintf(t.Format("20060102150405")+"%d",
		t.Nanosecond())
	return nowStr
}

// XqkTimeToStr13 格式化时间 => "2021年4月9日 12时4分46秒"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr13(t time.Time) string {
	nowStr := fmt.Sprintf("%d年%d月%d日 %d时%d分%d秒",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// XqkTimeToStr14 格式化时间 => "2021年04月09日 12时04分46秒"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr14(t time.Time) string {
	return t.Format("2006年01月02日 15时04分05秒")
}

// XqkTimeToStr15 格式化时间 => "2021年4月9日"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr15(t time.Time) string {
	nowStr := fmt.Sprintf("%d年%d月%d日",
		t.Year(),
		t.Month(),
		t.Day())
	return nowStr
}

// XqkTimeToStr16 格式化时间 => "2021年04月09日"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr16(t time.Time) string {
	return t.Format("2006年01月02日")
}

// XqkTimeToStr17 格式化时间 => "April 9(st), 2021"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr17(t time.Time) string {
	return fmt.Sprintf("%s %d(%s), %d",
		XqkTimeGetMonthEnglish(t.Month()),
		t.Day(),
		XqkIntGetEnglishNumTail(t.Day()),
		t.Year(),
	)
}

// XqkTimeToStr18 格式化时间 => "April 9, 2021"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr18(t time.Time) string {
	return fmt.Sprintf("%s %d, %d",
		XqkTimeGetMonthEnglish(t.Month()),
		t.Day(),
		t.Year(),
	)
}

// XqkTimeToStr19 格式化时间 => "April 9(st) 2021"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr19(t time.Time) string {
	return fmt.Sprintf("%s %d(%s) %d",
		XqkTimeGetMonthEnglish(t.Month()),
		t.Day(),
		XqkIntGetEnglishNumTail(t.Day()),
		t.Year(),
	)
}

// XqkTimeToStr20 格式化时间 => "April 9 2021"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr20(t time.Time) string {
	return fmt.Sprintf("%s %d %d",
		XqkTimeGetMonthEnglish(t.Month()),
		t.Day(),
		t.Year(),
	)
}

// XqkTimeToStr21 格式化时间 => "20210409"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr21(t time.Time) string {
	return t.Format("20060102")
}

// XqkTimeToStr22 格式化时间 => "120446"
// 其他格式参考 ToStrTransformMap
func XqkTimeToStr22(t time.Time) string {
	return t.Format("150405")
}
