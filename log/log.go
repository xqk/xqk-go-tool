package xqkLog

import xqkAll "github.com/xqk/xqk-go-tool/xqk_all"

// Error 打印错误日志（控制台显示红色），只接收一个变量，如需格式化请使用 LogErrorF
func Error(obj interface{}) {
	xqkAll.XqkLogError(obj)
}

// ErrorF 打印错误日志（控制台显示红色），参数如 fmt.Printf
func ErrorF(format string, a ...interface{}) {
	xqkAll.XqkLogErrorF(format, a...)
}

// ErrorFLn 打印错误日志（控制台显示红色）后面加一个换行，参数如 fmt.Printf
func ErrorFLn(format string, a ...interface{}) {
	xqkAll.XqkLogErrorFLn(format, a...)
}

// ErrorLn 在 LogError 后加一个换行
func ErrorLn(obj interface{}) {
	xqkAll.XqkLogErrorLn(obj)
}

// Success 打印错误日志（控制台显示绿色），只接收一个变量，如需格式化请使用 LogSuccessF
func Success(obj interface{}) {
	xqkAll.XqkLogSuccess(obj)
}

// SuccessF 打印错误日志（控制台显示绿色），参数如 fmt.Printf
func SuccessF(format string, a ...interface{}) {
	xqkAll.XqkLogSuccessF(format, a...)
}

// SuccessFLn 打印错误日志（控制台显示绿色）后面加一个换行，参数如 fmt.Printf
func SuccessFLn(format string, a ...interface{}) {
	xqkAll.XqkLogSuccessFLn(format, a...)
}

// SuccessLn 在 LogSuccess 后加一个换行
func SuccessLn(obj interface{}) {
	xqkAll.XqkLogSuccessLn(obj)
}

// Warning 打印错误日志（控制台显示黄色），只接收一个变量，如需格式化请使用 LogErrorF
func Warning(obj interface{}) {
	xqkAll.XqkLogWarning(obj)
}

// WarningF 打印错误日志（控制台显示黄色），参数如 fmt.Printf
func WarningF(format string, a ...interface{}) {
	xqkAll.XqkLogWarningF(format, a...)
}

// WarningFLn 打印错误日志（控制台显示黄色）后面加一个换行，参数如 fmt.Printf
func WarningFLn(format string, a ...interface{}) {
	xqkAll.XqkLogWarningFLn(format, a...)
}

// WarningLn 在 LogWarning 后加一个
func WarningLn(obj interface{}) {
	xqkAll.XqkLogWarningLn(obj)
}
