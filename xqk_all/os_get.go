package xqkAll

import (
	"os/exec"
	"runtime"
)

// XqkOsGetType 获取系统类型
// - android
// - darwin
// - dragonfly
// - freebsd
// - linux
// - nacl
// - netbsd
// - openbsd
// - plan9
// - solaris
// - windows
// 参考：https://github.com/golang/go/blob/master/src/go/build/syslist.go
func XqkOsGetType() string {
	return runtime.GOOS
}

// XqkOsGetGoarch 获取系统架构
// - 386
// - amd64
// - amd64p32
// - arm
// - arm64
// - ppc64
// - ppc64le
// - mips
// - mipsle
// - mips64
// - mips64le
// - mips64p32
// - mips64p32le
// - ppc
// - s390
// - s390x
// - sparc
// - sparc64
// 参考：https://github.com/golang/go/blob/master/src/go/build/syslist.go
func XqkOsGetGoarch() string {
	return runtime.GOARCH
}

// XqkOsGetCmd 获取命令行对象
// arg 中不要出现通配符，否则按字符串处理
func XqkOsGetCmd(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

// XqkOsGetCmdWithPrefix 获取命令行对象，带c前缀
func XqkOsGetCmdWithPrefix(cmd string) *exec.Cmd {
	if XqkOsIsTypeWindows() {
		return exec.Command("cmd", "/C", cmd)
	} else {
		return exec.Command("bash", "-c", cmd)
	}

}
