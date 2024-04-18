package xqkAll

import "os"

// XqkSStrGetMerge 获取合并字符串
// 缩写： XqkSStrGetM
//
// "a","b","c" >> "abc"
func XqkSStrGetMerge(sList ...string) string {
	result := ""
	for i := range sList {
		result += sList[i]
	}
	return result
}

// XqkSStrGetMergeByBlank 获取合并字符串，中间使用空格合并
// 缩写： XqkSStrGetMBB
//
// "a","b","c" >> "a b c"
func XqkSStrGetMergeByBlank(sList ...string) string {
	return XqkSStrGetMergeBySep(" ", sList...)
}

// XqkSStrGetMergeByOsPathSep 获取合并字符串，中间使用系统路径分隔符合并
func XqkSStrGetMergeByOsPathSep(sList ...string) string {
	return XqkSStrGetMergeBySep(string(os.PathSeparator), sList...)
}

// XqkSStrGetMergeBySep 获取合并字符串，指定中间的合并符
// 缩写： XqkSStrGetMBS
//
// "a","b","c" > "/" > "a/b/c"
func XqkSStrGetMergeBySep(Sep string, sList ...string) string {
	result := ""
	for i := range sList {
		if i == 0 {
			result += sList[i]
		} else {
			result += Sep + sList[i]
		}
	}
	return result
}

// XqkSStrGetM 获取合并字符串
// XqkSStrGetMerge 的缩写
//
// "a","b","c" >> "abc"
func XqkSStrGetM(sList ...string) string {
	return XqkSStrGetMerge(sList...)
}

// XqkSStrGetMBB 获取合并字符串，中间使用空格合并
// XqkSStrGetMergeByBlank 的缩写
//
// "a","b","c" >> "a b c"
func XqkSStrGetMBB(sList ...string) string {
	return XqkSStrGetMergeByBlank(sList...)
}

// XqkSStrGetMBS 获取合并字符串，指定中间的合并符
// XqkSStrGetMergeBySep 的缩写
//
// "a","b","c" > "/" > "a/b/c"
func XqkSStrGetMBS(sep string, sList ...string) string {
	return XqkSStrGetMergeBySep(sep, sList...)
}
