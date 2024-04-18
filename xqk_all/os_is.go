package xqkAll

// XqkOsIsTypeAndroid 是否是 android 系统
func XqkOsIsTypeAndroid() bool {
	return XqkOsGetType() == "android"
}

// XqkOsIsTypeDarwin 是否是 darwin 系统，比如Mac
func XqkOsIsTypeDarwin() bool {
	return XqkOsGetType() == "darwin"
}

// XqkOsIsTypeDragonfly 是否是 dragonfly 系统
func XqkOsIsTypeDragonfly() bool {
	return XqkOsGetType() == "dragonfly"
}

// XqkOsIsTypeFreebsd 是否是 freebsd 系统
func XqkOsIsTypeFreebsd() bool {
	return XqkOsGetType() == "freebsd"
}

// XqkOsIsTypeLinux 是否是 linux 系统
func XqkOsIsTypeLinux() bool {
	return XqkOsGetType() == "linux"
}

// XqkOsIsTypeNacl 是否是 nacl 系统
func XqkOsIsTypeNacl() bool {
	return XqkOsGetType() == "nacl"
}

// XqkOsIsTypeNetbsd 是否是 netbsd 系统
func XqkOsIsTypeNetbsd() bool {
	return XqkOsGetType() == "netbsd"
}

// XqkOsIsTypeOpenbsd 是否是 openbsd 系统
func XqkOsIsTypeOpenbsd() bool {
	return XqkOsGetType() == "openbsd"
}

// XqkOsIsTypePlan9 是否是 plan9 系统
func XqkOsIsTypePlan9() bool {
	return XqkOsGetType() == "plan9"
}

// XqkOsIsTypeSolaris 是否是 solaris 系统
func XqkOsIsTypeSolaris() bool {
	return XqkOsGetType() == "solaris"
}

// XqkOsIsTypeWindows 是否是 windows 系统
func XqkOsIsTypeWindows() bool {
	return XqkOsGetType() == "windows"
}

// ------------------------------------

// XqkOsIsGoarch386 是否是 386 体系架构
func XqkOsIsGoarch386() bool {
	return XqkOsGetGoarch() == "386"
}

// XqkOsIsGoarchAmd64 是否是 amd64 体系架构
func XqkOsIsGoarchAmd64() bool {
	return XqkOsGetGoarch() == "amd64"
}

// XqkOsIsGoarchAmd64p32 是否是 amd64p32 体系架构
func XqkOsIsGoarchAmd64p32() bool {
	return XqkOsGetGoarch() == "amd64p32"
}

// XqkOsIsGoarchArm 是否是 arm 体系架构
func XqkOsIsGoarchArm() bool {
	return XqkOsGetGoarch() == "arm"
}

// XqkOsIsGoarchArm64 是否是 arm64 体系架构
func XqkOsIsGoarchArm64() bool {
	return XqkOsGetGoarch() == "arm64"
}

// XqkOsIsGoarchPpc64 是否是 ppc64 体系架构
func XqkOsIsGoarchPpc64() bool {
	return XqkOsGetGoarch() == "ppc64"
}

// XqkOsIsGoarchPpc64le 是否是 ppc64le 体系架构
func XqkOsIsGoarchPpc64le() bool {
	return XqkOsGetGoarch() == "ppc64le"
}

// XqkOsIsGoarchMips 是否是 mips 体系架构
func XqkOsIsGoarchMips() bool {
	return XqkOsGetGoarch() == "mips"
}

// XqkOsIsGoarchMipsle 是否是 mipsle 体系架构
func XqkOsIsGoarchMipsle() bool {
	return XqkOsGetGoarch() == "mipsle"
}

// XqkOsIsGoarchMips64 是否是 mips64 体系架构
func XqkOsIsGoarchMips64() bool {
	return XqkOsGetGoarch() == "mips64"
}

// XqkOsIsGoarchMips64le 是否是 mips64le 体系架构
func XqkOsIsGoarchMips64le() bool {
	return XqkOsGetGoarch() == "mips64le"
}

// XqkOsIsGoarchMips64p32 是否是 mips64p32 体系架构
func XqkOsIsGoarchMips64p32() bool {
	return XqkOsGetGoarch() == "mips64p32"
}

// XqkOsIsGoarchMips64p32le 是否是 mips64p32le 体系架构
func XqkOsIsGoarchMips64p32le() bool {
	return XqkOsGetGoarch() == "mips64p32le"
}

// XqkOsIsGoarchPpc 是否是 ppc 体系架构
func XqkOsIsGoarchPpc() bool {
	return XqkOsGetGoarch() == "ppc"
}

// XqkOsIsGoarchS390 是否是 s390 体系架构
func XqkOsIsGoarchS390() bool {
	return XqkOsGetGoarch() == "s390"
}

// XqkOsIsGoarchS390x 是否是 s390x 体系架构
func XqkOsIsGoarchS390x() bool {
	return XqkOsGetGoarch() == "s390x"
}

// XqkOsIsGoarchSparc 是否是 sparc 体系架构
func XqkOsIsGoarchSparc() bool {
	return XqkOsGetGoarch() == "sparc"
}

// XqkOsIsGoarchSparc64 是否是 sparc64 体系架构
func XqkOsIsGoarchSparc64() bool {
	return XqkOsGetGoarch() == "sparc64"
}
