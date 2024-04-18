package xqkFile

import (
	"os"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func DoCopy(src, dest string) error {
	return xqkAll.XqkFileDoCopy(src, dest)
}

func DoCreate(filePath string) error {
	return xqkAll.XqkFileDoCreate(filePath)
}

// DoOpen 使用 os.O_RDWR, os.ModePerm 打开一个路径，
// 可读可写
// 如果文件不存在就会报错
// 如果需要其他默认请直接使用 os.OpenFile()
// 记得 file.Close()
func DoOpen(filePath string) (*os.File, error) {
	return xqkAll.XqkFileDoOpen(filePath)
}
