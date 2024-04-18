package xqkAll

import (
	"bytes"
	"os/exec"
	"path"
	"strings"

	xqkM "github.com/xqk/xqk-go-tool/xqk_model"
)

// XqkOsDoRunCmd 执行命令行，只返回 是否出错
func XqkOsDoRunCmd(name string, arg ...string) error {
	return XqkOsGetCmd(name, arg...).Run()
}

// XqkOsDoRunCmdWithResult 执行命令行，返回 执行结果和是否出错
func XqkOsDoRunCmdWithResult(name string, arg ...string) ([]byte, error) {
	return XqkOsGetCmd(name, arg...).CombinedOutput()
}

// XqkOsDoRunCmdWithOutAndErr 执行命令行，返回 错误结果、输出结果和是否出错
func XqkOsDoRunCmdWithOutAndErr(name string, arg ...string) ([]byte, []byte, error) {
	cmd := XqkOsGetCmd(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

// XqkOsDoCmdAddPipe 按顺序加管道
// 即：c1 | c2
func XqkOsDoCmdAddPipe(cList ...*exec.Cmd) error {
	if len(cList) <= 2 {
		return nil
	}
	var err error
	for i := range cList {
		if i == 0 {
			continue
		}
		cList[i].Stdin, err = cList[i-1].StdoutPipe()
		if err != nil {
			return err
		}
	}
	return err
}

// XqkOsDoRunCmdPipe 按顺序加管道执行命令
func XqkOsDoRunCmdPipe(cList []*exec.Cmd) error {
	if len(cList) <= 0 {
		return nil
	}
	err := XqkOsDoCmdAddPipe(cList...)
	if err != nil {
		return err
	}
	for i := range cList {
		if i == 0 {
			continue
		}
		err := cList[i].Start()
		if err != nil {
			return err
		}
	}
	err = cList[0].Run()
	if err != nil {
		return err
	}
	for i := range cList {
		if i == 0 {
			continue
		}
		err := cList[i].Wait()
		if err != nil {
			return err
		}
	}
	return err
}

// XqkOsDoBuildCmdPipe 自动组装cmd，按顺序加管道执行命令
func XqkOsDoBuildCmdPipe(cmdStrList []xqkM.CmdStr) error {
	var cList []*exec.Cmd
	for i := range cmdStrList {
		cList = append(cList, XqkOsGetCmd(cmdStrList[i].Name, cmdStrList[i].Arg...))
	}
	return XqkOsDoRunCmdPipe(cList)
}

// XqkOsDoRunCmdWidthPrefix 执行命令行，在前面加上c前缀
func XqkOsDoRunCmdWidthPrefix(cmd string) error {
	return XqkOsGetCmdWithPrefix(cmd).Run()
}

// XqkOsDoOpenFileManager 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件。
func XqkOsDoOpenFileManager(path string) error {
	if XqkOsIsTypeWindows() {
		err := XqkOsDoRunCmd("cmd", "/C", "explorer "+path)
		if err != nil {
			if XqkErrIsExitStatus1(err) {
				err = nil
			}
		}
		return err
	} else if XqkOsIsTypeDarwin() {
		// Mac
		err := XqkOsDoRunCmd("bash", "-c", "open "+path)
		if err != nil {
			if XqkErrIsExitStatus1(err) {
				err = nil
			}
		}
		return err
	} else if XqkOsIsTypeLinux() {
		// Linux
		err := XqkOsDoRunCmd("bash", "-c", "nautilus "+path)
		if err != nil {
			if XqkErrIsExitStatus1(err) {
				err = nil
			}
		}
		return err
	}
	return nil
}

// XqkOsDoOpenFileManagerByParent 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件的父文件。
func XqkOsDoOpenFileManagerByParent(p string) error {
	tp := strings.ReplaceAll(p, "\\", "/")
	if tp == p {
		return XqkOsDoOpenFileManager(path.Dir(p))
	} else {
		return XqkOsDoOpenFileManager(strings.ReplaceAll(path.Dir(tp), "/", "\\"))
	}
}
