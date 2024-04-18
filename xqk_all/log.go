package xqkAll

import "fmt"

// XqkLogError 打印错误日志（控制台显示红色），只接收一个变量，如需格式化请使用 LogErrorF
func XqkLogError(obj interface{}) {
	fmt.Printf("%c[1;0;31m%v%c[0m", 0x1B, obj, 0x1B)
}

// XqkLogErrorLn 在 LogError 后加一个换行
func XqkLogErrorLn(obj interface{}) {
	XqkLogErrorF("%v\n", obj)
}

// XqkLogErrorF 打印错误日志（控制台显示红色），参数如 fmt.Printf
func XqkLogErrorF(format string, a ...interface{}) {
	XqkLogError(fmt.Sprintf(format, a...))
}

// XqkLogErrorFLn 打印错误日志（控制台显示红色）后面加一个换行，参数如 fmt.Printf
func XqkLogErrorFLn(format string, a ...interface{}) {
	XqkLogErrorLn(fmt.Sprintf(format, a...))
}

// XqkLogSuccess 打印错误日志（控制台显示绿色），只接收一个变量，如需格式化请使用 LogSuccessF
func XqkLogSuccess(obj interface{}) {
	fmt.Printf("%c[1;0;32m%v%c[0m", 0x1B, obj, 0x1B)
}

// XqkLogSuccessLn 在 LogSuccess 后加一个换行
func XqkLogSuccessLn(obj interface{}) {
	XqkLogSuccessF("%v\n", obj)
}

// XqkLogSuccessF 打印错误日志（控制台显示绿色），参数如 fmt.Printf
func XqkLogSuccessF(format string, a ...interface{}) {
	XqkLogSuccess(fmt.Sprintf(format, a...))
}

// XqkLogSuccessFLn 打印错误日志（控制台显示绿色）后面加一个换行，参数如 fmt.Printf
func XqkLogSuccessFLn(format string, a ...interface{}) {
	XqkLogSuccessLn(fmt.Sprintf(format, a...))
}

// XqkLogWarning 打印错误日志（控制台显示黄色），只接收一个变量，如需格式化请使用 LogErrorF
func XqkLogWarning(obj interface{}) {
	fmt.Printf("%c[1;0;33m%v%c[0m", 0x1B, obj, 0x1B)
}

// XqkLogWarningLn 在 LogWarning 后加一个
func XqkLogWarningLn(obj interface{}) {
	XqkLogWarningF("%v\n", obj)
}

// XqkLogWarningF 打印错误日志（控制台显示黄色），参数如 fmt.Printf
func XqkLogWarningF(format string, a ...interface{}) {
	XqkLogWarning(fmt.Sprintf(format, a...))
}

// XqkLogWarningFLn 打印错误日志（控制台显示黄色）后面加一个换行，参数如 fmt.Printf
func XqkLogWarningFLn(format string, a ...interface{}) {
	XqkLogWarningLn(fmt.Sprintf(format, a...))
}
