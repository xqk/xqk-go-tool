package xqkAll

import (
	"io/fs"
	"io/ioutil"
	"os"

	xqkVar "github.com/xqk/xqk-go-tool/xqk_var"
)

// XqkDirDoMkDir 使用 os.ModePerm 模式创建一个文件夹，参数只能是一个文件夹名，不能是路径，
// 如果需要其他模式，直接调用 os.Mkdir()
func XqkDirDoMkDir(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}

// XqkDirDoMkDirAll 使用 os.ModePerm 模式递归创建目录，
// 如果需要其他模式，直接调用 os.MkdirAll()
func XqkDirDoMkDirAll(dirName string) error {
	return os.MkdirAll(dirName, os.ModePerm)
}

// XqkDirDoMkDirByTime 根据时间创建字符串创建一个文件夹，
// 时间字符串参考 XqkTimeGetNowStr8
func XqkDirDoMkDirByTime(dirName string) error {
	return os.Mkdir(dirName+string(os.PathSeparator)+XqkTimeGetNowStr8(), os.ModePerm)
}

// XqkDirDoMkDirAllByTime 根据时间归创创建字符串创建一个文件夹，
// 时间字符串参考 XqkTimeGetNowStr8
func XqkDirDoMkDirAllByTime(dirName string) error {
	return os.MkdirAll(dirName+string(os.PathSeparator)+XqkTimeGetNowStr8(), os.ModePerm)
}

// XqkDirDoCopy 拷贝文件夹(包括内容)
//
// Go中路径需要注意：
// "/"：标识根目录
// "./"：项目根目录
func XqkDirDoCopy(src, dest string) error {
	if !XqkDirIsExists(src) {
		return xqkVar.BaseErrDirNotExists
	}
	if XqkDirIsExists(dest) {
		return xqkVar.BaseErrDirExists
	}
	rd, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	err = XqkDirDoMkDirAll(dest)
	if err != nil {
		return err
	}
	for i := range rd {
		err = copyDirItem(src, dest, rd[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func copyDirItem(src, dest string, file fs.FileInfo) error {
	var err error
	srcFile := src + string(os.PathSeparator) + file.Name()
	destFile := dest + string(os.PathSeparator) + file.Name()
	if file.IsDir() {
		// 目录
		err = XqkDirDoMkDir(destFile)
		if err != nil {
			return err
		}
		rd, readErr := ioutil.ReadDir(srcFile)
		if readErr != nil {
			return readErr
		}
		for i := range rd {
			err = copyDirItem(srcFile, destFile, rd[i])
			if err != nil {
				return err
			}
		}
	} else {
		// 文件
		err = XqkFileDoCopy(srcFile, destFile)
		if err != nil {
			return err
		}
	}
	return err
}
