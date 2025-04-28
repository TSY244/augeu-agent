package engUtils

import (
	"augeu-agent/pkg/fileSystem"
	"augeu-agent/pkg/logger"
	"os"
)

type FileSysUtils struct {
}

// NewFileSys 创建文件系统工具实例
//
// return:
//   - *FileSysUtils 文件系统工具实例
func NewFileSys() *FileSysUtils {
	return &FileSysUtils{}
}

// LsFile 获取指定路径下的所有文件名
//
// params:
//   - path 路径
//
// return:
//   - []string 文件名列表，错误时返回nil
//
// notice:
//  1. 如果路径无效或发生错误，返回nil
func (f *FileSysUtils) LsFile(path string) []string {
	files, err := fileSystem.LsFile(path)
	if err != nil {
		logger.Errorf("get path file error: %v", err)
		return nil
	}
	return files
}

// LsDir 获取指定路径下的所有目录名
//
// params:
//   - path 路径
//
// return:
//   - []string 目录名列表，错误时返回nil
//
// notice:
//  1. 如果路径无效或发生错误，返回nil
func (f *FileSysUtils) LsDir(path string) []string {
	dirs, err := fileSystem.LsDir(path)
	if err != nil {
		logger.Errorf("get path dir error: %v", err)
		return nil
	}
	return dirs
}

// GetHashWithFilePath 获取指定文件的哈希值
//
// params:
//   - filePath 文件路径
//
// return:
//   - string 哈希值，错误时返回空字符串
//
// notice:
//  1. 如果文件不存在或计算哈希失败，返回空字符串
func (f *FileSysUtils) GetHashWithFilePath(filePath string) string {
	hash, err := fileSystem.GetHashWithFilePath(filePath)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return ""
	}
	return hash
}

// GetHashesWithFilePaths 批量获取文件哈希值
//
// params:
//   - filePaths 文件路径列表
//
// return:
//   - []string 哈希值列表，错误时返回nil
//
// notice:
//  1. 如果任意文件哈希计算失败，整体返回nil
func (f *FileSysUtils) GetHashesWithFilePaths(filePaths []string) []string {
	hashes, err := fileSystem.GetHashsWithFilePaths(filePaths)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return nil
	}
	return hashes
}

// IntoFile 创建文件并写入数据
//
// params:
//   - fileName 文件名
//   - data 写入的数据
//
// return:
//   - error 错误信息，成功时为nil
//
// notice:
//  1. 如果文件创建失败或写入失败，返回相应错误
func (f *FileSysUtils) IntoFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		logger.Errorf("write file error: %v", err)
		return err
	}
	return nil
}

// AddToFile 向文件追加数据
//
// params:
//   - fileName 文件名
//   - data 追加的数据
//
// return:
//   - error 错误信息，成功时为nil
//
// notice:
//  1. 如果文件打开失败或追加数据失败，返回相应错误
func (f *FileSysUtils) AddToFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		logger.Errorf("write file error: %v", err)
		return err
	}
	return nil
}

// StrSpliceIntoFile 将字符串切片逐行写入文件
//
// params:
//   - fileName 文件名
//   - strSlice 字符串切片
//
// notice:
//  1. 如果文件创建失败或写入失败，记录错误日志
//  2. 每个字符串会以换行符分隔写入文件
func (f *FileSysUtils) StrSpliceIntoFile(fileName string, strSlice []string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return
	}
	defer file.Close()
	for _, str := range strSlice {
		_, err = file.WriteString(str + "\n")
		if err != nil {
			logger.Errorf("write file error: %v", err)
		}
	}
}
