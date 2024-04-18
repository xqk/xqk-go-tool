package xqkGen

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func Gen(gitPath string, moduleF func() []XqkModule, methodF func() map[string]map[string][]string) {
	allModule := moduleF()
	// allModule := []XqkModule{
	// 	{
	// 		PackageName: "xqkBool",
	// 		DirName:     "bool",
	// 		MethodModule: []XqkMethodModule{
	// 			{
	// 				Type:     To,
	// 				FileName: "bool_to.go",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		PackageName: "xqkLog",
	// 		DirName:     "log",
	// 		MethodModule: []XqkMethodModule{
	// 			{
	// 				Type:     NoType,
	// 				FileName: "log.go",
	// 			},
	// 		},
	// 	},
	// }
	methodMap := methodF()
	for i := range allModule {
		packageName := allModule[i].PackageName
		xqkAll.XqkLogSuccessLn(xqkAll.XqkIntToStr(i+1) + "." + packageName)
		err := genPackage(gitPath, allModule[i], methodMap, moduleF)
		if err != nil {
			return
		}
		printSep()
	}
}

func genPackage(gitPath string, p XqkModule, methodMap map[string]map[string][]string, moduleF func() []XqkModule) error {
	packageName := p.PackageName
	dirName := p.DirName

	// 1.检查`method_map.go`是否定义了`packageName`的方法
	xqkAll.XqkLogSuccessLn("\t◎ 检查[S]：`method_map.go`是否为`" + packageName + "`定义了方法。")
	if len(methodMap[packageName]) == 0 {
		xqkAll.XqkLogErrorLn("\t◎ 结果[E]：`method_map.go`未给`" + packageName + "`定义了方法，请检查`method_map.go`!")
		return errors.New("")
	}
	xqkAll.XqkLogSuccessLn("\t◎ 结果[S]：`method_map.go`为`" + packageName + "`定义了方法。\n")

	// 2.检查文件夹
	xqkAll.XqkLogSuccessLn("\t◎ 检查[S]：`" + dirName + "`文件夹。")
	if xqkAll.XqkDirIsExists(dirName) {
		xqkAll.XqkLogSuccessLn("\t◎ 结果[S]：`" + dirName + "`文件夹已存在。\n")
	} else {
		xqkAll.XqkLogSuccessLn("\t◎ 结果[S]：`" + dirName + "`文件夹不存在。\n")
		xqkAll.XqkLogSuccessLn("\t◎ 操作[S]：创建`" + dirName + "`文件夹")
		err := xqkAll.XqkDirDoMkDir(dirName)
		if err != nil {
			xqkAll.XqkLogErrorLn(err.Error())
			xqkAll.XqkLogErrorLn("\t◎ 结果[E]：创建`" + dirName + "`文件夹失败!\n")
			return err
		}
		xqkAll.XqkLogSuccessLn("\t◎ 结果[S]：创建`" + dirName + "`文件夹成功!\n")
	}

	// 3.检查`module_list.go`是否为`packageName`定义了模块
	xqkAll.XqkLogSuccessLn("\t◎ 检查[S]：`module_list.go`是否为`" + packageName + "`定义了模块。")
	if len(p.MethodModule) == 0 {
		xqkAll.XqkLogErrorLn("\t◎ 结果[E]：`module_list.go`未给`" + packageName + "`定义了模块，请检查`module_list.go`!")
		return errors.New("")
	}
	xqkAll.XqkLogSuccessLn("\t◎ 结果[S]：`method_map.go`为`" + packageName + "`定义了模块。\n")

	// 4.处理`packageName`下的模块
	for i := range p.MethodModule {
		moduleTypeStr := string(p.MethodModule[i].Type)
		moduleFileName := p.MethodModule[i].FileName
		srcFilePath := xqkAll.XqkSStrGetMergeByOsPathSep("xqk_all", moduleFileName)
		genFilePath := xqkAll.XqkSStrGetMergeByOsPathSep(dirName, moduleFileName)
		// 4.1.检查模块类型是否存在
		if moduleTypeStr == "" {
			xqkAll.XqkLogErrorLn("\t- [E]`" + packageName + "`模块类型未定义，请检查`module_list.go`!")
			return errors.New("")
		}
		xqkAll.XqkLogSuccessLn("\t- [S]`" + packageName + "`-`" + moduleTypeStr + "`")

		// 4.2.检查文件是否定义
		xqkAll.XqkLogSuccessLn("\t\t◎ 检查[S]：`" + packageName + "`模块的定义文件。")
		if p.MethodModule[i].FileName == "" {
			xqkAll.XqkLogErrorLn("\t\t◎ 结果[E]：`" + packageName + "`模块类型未定义，请检查`module_list.go`!")
			return errors.New("")
		}
		xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：`" + packageName + "`模块已定义文件。\n")

		// 4.3.检查定义的源文件是否存在
		xqkAll.XqkLogSuccessLn("\t\t◎ 检查[S]：`" + packageName + "`模块的定义文件对应的源文件是否存在。")
		if !xqkAll.XqkFileIsExists(srcFilePath) {
			xqkAll.XqkLogErrorLn("\t\t◎ 结果[E]：`" + srcFilePath + "`文件不存在。")
			xqkAll.XqkLogErrorLn("\t\t◎         `" + packageName + "`的`" + moduleTypeStr + "`对应的`" + moduleFileName + "`文件不存在!请检查`module_list.go`!")
			return errors.New("")
		}
		xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：`" + packageName + "`模块的定义文件对应的源文件存在。\n")

		// 4.4.判断`methodMap`只是否定义了该模块类型
		xqkAll.XqkLogSuccessLn("\t\t◎ 检查[S]：`method_map.go`中是否定义了`" + packageName + "`的`" + string(p.MethodModule[i].Type) + "`类型的方法。")
		if len(methodMap[packageName][string(p.MethodModule[i].Type)]) == 0 {
			xqkAll.XqkLogErrorLn("\t\t◎ 结果[E]：`method_map.go`中未定义了`" + packageName + "`的`" + string(p.MethodModule[i].Type) + "`类型的方法!请检查`method_map.go`。")
			return errors.New("")
		}
		xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：`method_map.go`定义了`" + packageName + "`的`" + string(p.MethodModule[i].Type) + "`类型的方法。\n")

		// 4.5.生成输出文件
		xqkAll.XqkLogSuccessLn("\t\t◎ 操作[S]：生成`" + genFilePath + "`文件。")
		err := genFile(
			gitPath,
			p.DirName,
			moduleFileName,
			srcFilePath,
			packageName,
			string(p.MethodModule[i].Type),
			methodMap)
		if err != nil {
			xqkAll.XqkLogErrorLn("\t\t◎ 结果[E]：生成`" + genFilePath + "`文件失败!")
			return err
		}
		xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：生成`" + genFilePath + "`文件成功。\n")

		// 4.6.检查是否有测试文件
		testFilePath := xqkAll.XqkStrGetTrimRightSStr(srcFilePath, ".go") + "_test.go"
		xqkAll.XqkLogSuccessLn("\t\t◎ 检查[S]：是否存在`" + testFilePath + "`测试文件。")
		if !xqkAll.XqkFileIsExists(testFilePath) {
			xqkAll.XqkLogWarningLn("\t\t◎ 结果[W]：不存在`" + testFilePath + "`测试文件。\n")
		} else {
			xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：存在`" + testFilePath + "`测试文件。\n")

			// 生成测试文件
			genTestFilePath := xqkAll.XqkStrGetTrimRightSStr(genFilePath, ".go") + "_test.go"
			xqkAll.XqkLogSuccessLn("\t\t◎ 操作[S]：生成`" + genTestFilePath + "`测试文件。")
			err = genTestFile(
				gitPath,
				dirName,
				xqkAll.XqkStrGetTrimRightSStr(moduleFileName, ".go")+"_test.go",
				testFilePath,
				packageName,
			)
			if err != nil {
				xqkAll.XqkLogErrorLn("\t\t◎ 结果[E]：生成`" + genTestFilePath + "`测试文件失败!")
				return err
			}
			xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：生成`" + genTestFilePath + "`测试文件成功。\n")
		}

		// 4.7.检查是否有案例文件
		exampleFilePath := xqkAll.XqkStrGetTrimRightSStr(srcFilePath, ".go") + "_example_test.go"
		xqkAll.XqkLogSuccessLn("\t\t◎ 检查[S]：是否存在`" + exampleFilePath + "`案例文件。")
		if !xqkAll.XqkFileIsExists(exampleFilePath) {
			xqkAll.XqkLogWarningLn("\t\t◎ 结果[W]：不存在`" + exampleFilePath + "`案例文件。\n")
		} else {
			xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：存在`" + exampleFilePath + "`案例文件。\n")
			// 生成案例文件
			genExampleFilePath := xqkAll.XqkStrGetTrimRightSStr(exampleFilePath, ".go") + "_test.go"
			xqkAll.XqkLogSuccessLn("\t\t◎ 操作[S]：生成`" + genExampleFilePath + "`案例文件。")
			err = genExampleFile(
				gitPath,
				dirName,
				xqkAll.XqkStrGetTrimRightSStr(moduleFileName, ".go")+"_example_test.go",
				exampleFilePath,
				packageName+"_test",
				moduleF,
			)
			if err != nil {
				xqkAll.XqkLogErrorLn("\t\t◎ 结果[E]：生成`" + genExampleFilePath + "`案例文件失败!")
				return err
			}
			xqkAll.XqkLogSuccessLn("\t\t◎ 结果[S]：生成`" + genExampleFilePath + "`案例文件成功。\n")
		}
	}
	return nil
}

func genExampleFile(gitPath, dirName, fileName, srcFilePath, packageName string, moduleF func() []XqkModule) error {
	pathStr := xqkAll.XqkSStrGetMergeByOsPathSep(dirName, fileName)
	// 1.创建文件
	// p.DirName / p.MethodModule[i].FileName
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：创建`" + pathStr + "`文件。")
	file, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：创建`" + pathStr + "`文件失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：创建`" + pathStr + "`文件成功。\n")
	// 2.写入包名
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入包名`" + packageName + "`。")
	_, err = writer.WriteString("package " + packageName + "\r\n\r\n")
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入包名`" + packageName + "`失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入包名`" + packageName + "`成功。\n")

	// 3.读取测试文件
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`测试文件内容。")
	content, err := genExampleFileStr(gitPath, srcFilePath, packageName, moduleF)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`测试文件内容失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`测试文件内容成功。\n")

	// 4.写入
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：测试内容写入`" + pathStr + "`文件。")
	_, err = writer.WriteString(content)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：测试内容写入`" + pathStr + "`文件失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：测试内容写入`" + pathStr + "`文件成功。\n")

	// 保存文件
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：保存`" + pathStr + "`文件的写入内容。")
	err = writer.Flush()
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：保存`" + pathStr + "`文件的写入内容失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：保存`" + pathStr + "`文件的写入内容成功。\n")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}

func genExampleFileStr(gitPath, filePath, packageName string, moduleF func() []XqkModule) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	codeStr := ""
	inImport := false
	var addImport []string
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if strings.HasPrefix(str, "package ") {
			continue
		}
		if strings.HasPrefix(str, "import (") {
			inImport = true
			continue
		}
		if strings.HasPrefix(str, "import ") {
			continue
		}
		if inImport {
			if strings.HasPrefix(str, ")") {
				inImport = false
			}
			continue
		}
		ts, addImportStr, eErr := handleExampleLine(gitPath, str, packageName, moduleF)
		if eErr != nil {
			return "", eErr
		}
		if addImportStr != "" {
			addImport = append(addImport, addImportStr)
		}
		codeStr += ts
	}
	importStr, err := getImportStr(filePath, false, addImport...)
	if err != nil {
		return outStr, err
	}
	importStr = strings.ReplaceAll(importStr, "\txqkAll \""+gitPath+"/xqk_all\"\r\n", "")
	importStr = strings.ReplaceAll(importStr, "\txqkAll \""+gitPath+"/xqk_all\"\n", "")
	importStr = strings.ReplaceAll(importStr, "\txqkAll \""+gitPath+"/xqk_all\"", "")
	return importStr + "\r\n" + xqkAll.XqkStrGetTrimSStr(codeStr, "\r", "\n") + "\r\n", nil
}

func genTestFile(gitPath, dirName, fileName, srcFilePath, packageName string) error {
	pathStr := xqkAll.XqkSStrGetMergeByOsPathSep(dirName, fileName)
	// 1.创建文件
	// p.DirName / p.MethodModule[i].FileName
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：创建`" + pathStr + "`文件。")
	file, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：创建`" + pathStr + "`文件失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：创建`" + pathStr + "`文件成功。\n")
	// 2.写入包名
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入包名`" + packageName + "`。")
	_, err = writer.WriteString("package " + packageName + "\r\n\r\n")
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入包名`" + packageName + "`失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入包名`" + packageName + "`成功。\n")

	// 3.读取测试文件
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`测试文件内容。")
	content, err := getTestFileStr(gitPath, srcFilePath, packageName)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`测试文件内容失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`测试文件内容成功。\n")

	// 4.写入
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：测试内容写入`" + pathStr + "`文件。")
	_, err = writer.WriteString(content)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：测试内容写入`" + pathStr + "`文件失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：测试内容写入`" + pathStr + "`文件成功。\n")

	// 保存文件
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：保存`" + pathStr + "`文件的写入内容。")
	err = writer.Flush()
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：保存`" + pathStr + "`文件的写入内容失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：保存`" + pathStr + "`文件的写入内容成功。\n")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}

func genFile(gitPath, dirName, fileName, srcFilePath, packageName, methodType string, methodMap map[string]map[string][]string) error {
	pathStr := xqkAll.XqkSStrGetMergeByOsPathSep(dirName, fileName)
	// 1.创建文件
	// p.DirName / p.MethodModule[i].FileName
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：创建`" + pathStr + "`文件。")
	file, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：创建`" + pathStr + "`文件失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：创建`" + pathStr + "`文件成功。\n")

	// 2.写入包名
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入包名`" + packageName + "`。")
	_, err = writer.WriteString("package " + packageName + "\r\n\r\n")
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入包名`" + packageName + "`失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入包名`" + packageName + "`成功。\n")

	// 3.写入导入包xqkAll
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`文件的导入包。")
	importStr, err := getImportStr(srcFilePath, true, "xqkAll \""+gitPath+"/xqk_all\"")
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：`" + srcFilePath + "`文件的导入包失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`文件的导入包成功。\n")

	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入导入包。")
	// _, err = writer.WriteString("import xqkAll \"github.com/xqk/xqk-go-tool/xqk_all\"" + "\r\n")
	_, err = writer.WriteString(importStr)
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入\"xqkAll\"导入包失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入\"xqkAll\"导入包成功。\n")

	// 4.写入方法
	methodList := methodMap[packageName][methodType]
	for i := range methodList {
		srcMethodName := strings.Replace(packageName+methodList[i], "y", "Y", 1)
		xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法。")
		methodStr, mErr := getMethodSTr(srcFilePath, methodList[i], srcMethodName, packageName)
		if mErr != nil {
			xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法失败!")
			xqkAll.XqkLogErrorLn(mErr.Error())
			return mErr
		}
		if methodStr == "" {
			xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法结果为空!")
			return errors.New("")
		} else {
			xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法结果成功。\n")
		}

		xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入`" + methodList[i] + "`方法。")
		_, err = writer.WriteString("\r\n" + methodStr)
		if err != nil {
			xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入`" + methodList[i] + "`方法失败!")
			xqkAll.XqkLogErrorLn(err.Error())
			return err
		}
		xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入`" + methodList[i] + "`方法成功。\n")
	}

	// 保存文件
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 操作[S]：保存`" + pathStr + "`文件的写入内容。")
	err = writer.Flush()
	if err != nil {
		xqkAll.XqkLogErrorLn("\t\t\t◎ 结果[E]：保存`" + pathStr + "`文件的写入内容失败!")
		xqkAll.XqkLogErrorLn(err.Error())
		return err
	}
	xqkAll.XqkLogSuccessLn("\t\t\t◎ 结果[S]：保存`" + pathStr + "`文件的写入内容成功。\n")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}

func getMethodSTr(filePath, methodName, srcMethodName, packageName string) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	commentStr := ""
	methodLine := ""
	inMethod := false
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if inMethod {
			if strings.HasPrefix(str, "}") {
				// 方法结尾了
				// 从methodLine中解析出参数来
				hasR, rErr := hasReturn(methodLine)
				if rErr != nil {
					return "", rErr
				}
				parameterList, hasDot, mErr := getMethodParameterList(methodLine)
				if mErr != nil {
					return "", mErr
				}
				outStr += "\t"
				if hasR {
					outStr += "return "
				}
				outStr += "xqkAll." + srcMethodName + "("
				for i := range parameterList {
					if i+1 == len(parameterList) {
						if hasDot {
							outStr += parameterList[i] + "..."
						} else {
							outStr += parameterList[i]
						}
					} else {
						outStr += parameterList[i] + ", "
					}
				}
				outStr += ")" + "\r\n"
				outStr += str
				break
			}
		}
		if strings.HasPrefix(str, "//") {
			// 如果是注释
			commentStr += str
		} else {
			// 不是注释，看是不是我们需要的方法
			if strings.HasPrefix(str, "func "+srcMethodName+"(") {
				// 是我们需要的方法
				inMethod = true
				methodLine = str
				commentStr = strings.ReplaceAll(commentStr, srcMethodName, methodName)
				comment, cErr := handleComment(commentStr, packageName)
				if cErr != nil {
					return "", cErr
				}
				outStr += comment
				outStr += strings.ReplaceAll(str, "func "+srcMethodName+"(", "func "+methodName+"(")
			} else {
				// 不是我们需要的方法
				commentStr = ""
			}
		}
	}
	return outStr, nil
}

func handleComment(s, packageName string) (string, error) {
	packageName = strings.Replace(packageName, "y", "Y", 1)
	// 已包名开头的就是当前包的方法，要删除该前缀
	selfPackageReg := regexp.MustCompile(packageName + `(Get|Op|Is|Do|To)`)
	if selfPackageReg == nil {
		return "", errors.New("handleComment-selfPackageReg-构造正则失败")
	}
	s = selfPackageReg.ReplaceAllStringFunc(s, func(ts string) string {
		return xqkAll.XqkStrGetTrimLeftSStr(ts, packageName)
	})
	funcReg := regexp.MustCompile(`Xqk[A-Z][\w]+(Get|Op|Is|Do|To)`)
	if funcReg == nil {
		return "", errors.New("handleComment-funcReg-构造正则失败")
	}
	s = funcReg.ReplaceAllStringFunc(s, func(ts string) string {
		ts = strings.Replace(ts, "Y", "y", 1)
		if strings.HasSuffix(ts, "Get") {
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Get") + ".Get"
		}
		if strings.HasSuffix(ts, "Op") {
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Op") + ".Op"
		}
		if strings.HasSuffix(ts, "Is") {
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Is") + ".Is"
		}
		if strings.HasSuffix(ts, "Do") {
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Do") + ".Do"
		}
		if strings.HasSuffix(ts, "To") {
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "To") + ".To"
		}
		return ts
	})
	s = strings.ReplaceAll(s, "XqkLog", "xqkLog.")
	return s, nil
}

func handleTestLine(s, packageName string) (string, bool, error) {
	hasxqkAll := false
	packageName = strings.Replace(packageName, "y", "Y", 1)
	// 已包名开头的就是当前包的方法，要删除该前缀
	selfPackageReg := regexp.MustCompile(packageName + `(Get|Op|Is|Do|To)`)
	if selfPackageReg == nil {
		return "", false, errors.New("handleTestLine-selfPackageReg-构造正则失败")
	}
	s = selfPackageReg.ReplaceAllStringFunc(s, func(ts string) string {
		return xqkAll.XqkStrGetTrimLeftSStr(ts, packageName)
	})
	funcReg := regexp.MustCompile(`Xqk[A-Z][\w]+(Get|Op|Is|Do|To)`)
	if funcReg == nil {
		return "", false, errors.New("handleTestLine-funcReg-构造正则失败")
	}
	s = funcReg.ReplaceAllStringFunc(s, func(ts string) string {
		hasxqkAll = true
		return "xqkAll." + ts
	})
	if strings.Contains(s, "XqkLog") {
		hasxqkAll = true
	}
	s = strings.ReplaceAll(s, "XqkLog", "xqkAll.XqkLog")
	return s, hasxqkAll, nil
}

func handleExampleLine(gitPath, s, packageName string, moduleF func() []XqkModule) (string, string, error) {
	funcReg := regexp.MustCompile(`xqkAll\.Xqk[A-Z][\w]+(Get|Op|Is|Do|To)`)
	if funcReg == nil {
		return "", "", errors.New("handleExampleLine-funcReg-构造正则失败")
	}
	linePackage := ""
	s = funcReg.ReplaceAllStringFunc(s, func(ts string) string {
		ts = strings.ReplaceAll(ts, "xqkAll.", "")
		suffixStr := ""
		if strings.HasSuffix(ts, "Get") {
			suffixStr = "Get"
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Get")
		}
		if strings.HasSuffix(ts, "Op") {
			suffixStr = "Op"
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Op")
		}
		if strings.HasSuffix(ts, "Is") {
			suffixStr = "Is"
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Is")
		}
		if strings.HasSuffix(ts, "Do") {
			suffixStr = "Do"
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "Do")
		}
		if strings.HasSuffix(ts, "To") {
			suffixStr = "To"
			ts = xqkAll.XqkStrGetTrimRightSStr(ts, "To")
		}
		linePackage = strings.Replace(ts, "Y", "y", 1)
		return linePackage + "." + suffixStr
	})
	s = strings.ReplaceAll(
		s,
		"Example"+strings.Replace(xqkAll.XqkStrGetTrimRightSStr(packageName, "_test"), "y", "Y", 1),
		"Example",
	)
	if linePackage == "" {
		return s, "", nil
	} else {
		packagePath := ""
		mList := moduleF()
		for i := range mList {
			if mList[i].PackageName == linePackage {
				packagePath = mList[i].DirName
				break
			}
		}
		if packagePath == "" {
			return "", "", errors.New("`module_list.go`中没有`" + linePackage + "`")
		}
		return s, linePackage + " \"" + gitPath + "/" + packagePath + "\"", nil
	}
}

func hasReturn(s string) (bool, error) {
	s = xqkAll.XqkStrGetDelEndRNStr(s)
	firstReverseBracketsIndex := -1
	bracketsNum := 1
	hasCheckFirstB := false
	for i := range s {
		if s[i] == '(' {
			if !hasCheckFirstB {
				hasCheckFirstB = true
			} else {
				bracketsNum += 1
			}
		}
		if s[i] == ')' {
			bracketsNum -= 1
		}
		if bracketsNum == 0 {
			firstReverseBracketsIndex = i
			break
		}
	}
	if firstReverseBracketsIndex == -1 {
		return false, errors.New("方法行检查返回值出错")
	}
	s = s[firstReverseBracketsIndex+1:]
	return xqkAll.XqkStrGetTrimSStr(s, " ", "{") != "", nil
}

func getMethodParameterList(s string) ([]string, bool, error) {
	startIndex := -1
	endIndex := -1
	bracketsNum := 1
	hasCheckFirstB := false
	for i := range s {
		if s[i] == '(' {
			if !hasCheckFirstB {
				hasCheckFirstB = true
				startIndex = i
			} else {
				bracketsNum++
			}
		}
		if s[i] == ')' {
			bracketsNum--
		}
		if bracketsNum == 0 {
			endIndex = i
			break
		}
	}
	if startIndex == -1 || endIndex == -1 || startIndex+1 > endIndex {
		return nil, false, errors.New("方法行检查参数列表出错")
	}
	s = xqkAll.XqkStrGetTrimSStr(s[startIndex+1:endIndex], " ")
	if s == "" {
		return nil, false, nil
	}
	var result []string
	hasDot := false
	var ts []byte
	bracketsNum = 0
	for i := range s {
		if s[i] == '(' {
			bracketsNum++
		}
		if s[i] == ')' {
			bracketsNum--
		}
		if !(s[i] == ',' && bracketsNum == 0) {
			ts = append(ts, s[i])
		}
		if (s[i] == ',' && bracketsNum == 0) || i+1 == len(s) {
			tts := xqkAll.XqkStrGetTrimSStr(string(ts), " ")
			blankIndex := strings.Index(tts, " ")
			if blankIndex == -1 {
				result = append(result, tts)
			} else {
				result = append(result, tts[:strings.Index(tts, " ")])
			}
			if i+1 == len(s) {
				hasDot = strings.Contains(tts[strings.Index(tts, " "):], "...")
			}
			ts = []byte{}
		}
	}
	return result, hasDot, nil
}

func getImportStr(filePath string, onlyMethod bool, addImport ...string) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	inBrackets := false
	var importStrList []string
	var methodImportStrList []string
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if strings.HasPrefix(str, "func Xqk") {
			tsList := xqkAll.XqkStrToStrListBySep(strings.ReplaceAll(str, "...", ""), " ")
			for i := range tsList {
				if strings.Contains(tsList[i], ".") {
					methodImportStrList = append(
						methodImportStrList,
						xqkAll.XqkStrGetTrimSStr(tsList[i][:strings.Index(tsList[i], ".")], "*", "[]", "("),
					)
				}
			}
		}
		if inBrackets {
			if strings.HasPrefix(str, ")") {
				inBrackets = false
			} else {
				importStrList = append(importStrList, xqkAll.XqkStrGetTrimSStr(xqkAll.XqkStrGetDelEndRNStr(str), "\t"))
			}
		} else {
			if strings.HasPrefix(str, "import (") {
				inBrackets = true
				continue
			}
			if strings.HasPrefix(str, "import ") {
				importStrList = append(importStrList, xqkAll.XqkStrGetTrimLeftSStr(xqkAll.XqkStrGetDelEndRNStr(str), "import "))
			}
		}
	}
	// 如果仅方法包
	if onlyMethod {
		xqkAll.XqkStrListOpFilter(&importStrList, func(x string) bool {
			packageName := ""
			if strings.HasPrefix(x, "\"") {
				// 冒号包
				ts := xqkAll.XqkStrGetTrimSStr(x, "\"")
				slashIndex := strings.LastIndex(ts, "/")
				if slashIndex == -1 {
					packageName = ts
				} else {
					packageName = ts[slashIndex+1:]
				}
			} else {
				// 重命名包
				packageName = x[:strings.Index(x, " ")]
			}
			for i := range methodImportStrList {
				if packageName == methodImportStrList[i] {
					return true
				}
			}
			return false
		})
	}

	for i := range addImport {
		importStrList = append(importStrList, addImport[i])
	}
	xqkAll.XqkStrListOpDeduplicate(&importStrList)
	if len(importStrList) == 1 {
		outStr = "import " + importStrList[0] + "\r\n"
	}
	if len(importStrList) > 1 {
		sort.Slice(importStrList, func(i, j int) bool {
			if importStrList[i] == importStrList[j] {
				return true
			}
			tsI := importStrList[i]
			tsJ := importStrList[j]
			iColon := strings.HasPrefix(importStrList[i], "\"")
			jColon := strings.HasPrefix(importStrList[j], "\"")
			if !iColon {
				tsI = tsI[strings.Index(tsI, " ")+1:]
			}
			if !jColon {
				tsJ = tsI[strings.Index(tsJ, " ")+1:]
			}
			return xqkAll.XqkRuneListIsLeByUnicodeAndLowerBeforeUpper(xqkAll.XqkStrToRuneList(tsI), xqkAll.XqkStrToRuneList(tsJ))
		})
		outStr += "import (\r\n"
		for i := range importStrList {
			outStr += "\t" + importStrList[i] + "\r\n"
		}
		outStr += ")\r\n"
	}
	return outStr, nil
}

func getTestFileStr(gitPath, filePath, packageName string) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	codeStr := ""
	inImport := false
	hasxqkAll := false
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if strings.HasPrefix(str, "package ") {
			continue
		}
		if strings.HasPrefix(str, "import (") {
			inImport = true
			continue
		}
		if strings.HasPrefix(str, "import ") {
			continue
		}
		if inImport {
			if strings.HasPrefix(str, ")") {
				inImport = false
			}
			continue
		}
		ts, tHasxqkAll, lErr := handleTestLine(str, packageName)
		if lErr != nil {
			return "", lErr
		}
		if !hasxqkAll {
			hasxqkAll = tHasxqkAll
		}
		codeStr += ts
	}
	var addImport []string
	if hasxqkAll {
		addImport = append(addImport, "xqkAll \""+gitPath+"/xqk_all\"")
	}
	importStr, err := getImportStr(filePath, false, addImport...)
	if err != nil {
		return outStr, err
	}
	if importStr == "" {
		return xqkAll.XqkStrGetTrimSStr(codeStr, "\r", "\n") + "\r\n", nil
	} else {
		return importStr + "\r\n" + xqkAll.XqkStrGetTrimSStr(codeStr, "\r", "\n") + "\r\n", nil
	}
}

func printSep() {
	xqkAll.XqkLogSuccess("\n")
	xqkAll.XqkLogSuccessLn("===========================")
	xqkAll.XqkLogSuccess("\n")
	xqkAll.XqkLogSuccess("\n")
}
