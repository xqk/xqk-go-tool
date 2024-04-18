package xqkAll

import (
	"io"
	"os"
	"path/filepath"

	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

// XqkFileDoOpen 使用 os.O_RDWR, os.ModePerm 打开一个路径，
// 可读可写
// 如果文件不存在就会报错
// 如果需要其他默认请直接使用 os.OpenFile()
// 记得 file.Close()
func XqkFileDoOpen(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
}

func XqkFileDoCreate(filePath string) error {
	if XqkFileIsExists(filePath) {
		return xqkVar.BaseErrFileExists
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return err
}

func XqkFileDoCopy(src, dest string) error {
	if !XqkFileIsExists(src) {
		return xqkVar.BaseErrFileNotExists
	}
	if XqkFileIsExists(dest) {
		return xqkVar.BaseErrFileExists
	}
	dir, _ := filepath.Split(dest)
	if dir != "" {
		err := XqkDirDoMkDirAll(dir)
		if err != nil {
			return err
		}
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)

	dstFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		_ = dstFile.Close()
	}(dstFile)

	_, err = io.Copy(dstFile, srcFile)
	return err
}
