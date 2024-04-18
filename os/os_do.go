package xqkOs

import (
	"os/exec"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
	xqkM "github.com/xqk/xqk-go-tool/xqk_model"
)

// DoBuildCmdPipe 自动组装cmd，按顺序加管道执行命令
func DoBuildCmdPipe(cmdStrList []xqkM.CmdStr) error {
	return xqkAll.XqkOsDoBuildCmdPipe(cmdStrList)
}

// DoCmdAddPipe 按顺序加管道
// 即：c1 | c2
func DoCmdAddPipe(cList ...*exec.Cmd) error {
	return xqkAll.XqkOsDoCmdAddPipe(cList...)
}

// DoOpenFileManager 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件。
func DoOpenFileManager(path string) error {
	return xqkAll.XqkOsDoOpenFileManager(path)
}

// DoOpenFileManagerByParent 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件的父文件。
func DoOpenFileManagerByParent(p string) error {
	return xqkAll.XqkOsDoOpenFileManagerByParent(p)
}

// DoRunCmd 执行命令行，只返回 是否出错
func DoRunCmd(name string, arg ...string) error {
	return xqkAll.XqkOsDoRunCmd(name, arg...)
}

// DoRunCmdPipe 按顺序加管道执行命令
func DoRunCmdPipe(cList []*exec.Cmd) error {
	return xqkAll.XqkOsDoRunCmdPipe(cList)
}

// DoRunCmdWidthPrefix 执行命令行，在前面加上c前缀
func DoRunCmdWidthPrefix(cmd string) error {
	return xqkAll.XqkOsDoRunCmdWidthPrefix(cmd)
}

// DoRunCmdWithOutAndErr 执行命令行，返回 错误结果、输出结果和是否出错
func DoRunCmdWithOutAndErr(name string, arg ...string) ([]byte, []byte, error) {
	return xqkAll.XqkOsDoRunCmdWithOutAndErr(name, arg...)
}

// DoRunCmdWithResult 执行命令行，返回 执行结果和是否出错
func DoRunCmdWithResult(name string, arg ...string) ([]byte, error) {
	return xqkAll.XqkOsDoRunCmdWithResult(name, arg...)
}
