package engUtils

import (
	"augeu-agent/pkg/fileSystem"
	"augeu-agent/pkg/logger"
	"os"
)

type FileSysUtils struct {
}

func NewFileSys() *FileSysUtils {
	return &FileSysUtils{}
}

func (f *FileSysUtils) LsFile(path string) []string {
	files, err := fileSystem.LsFile(path)
	if err != nil {
		logger.Errorf("get path file error: %v", err)
		return nil
	}
	return files
}

func (f *FileSysUtils) LsDir(path string) []string {
	dirs, err := fileSystem.LsDir(path)
	if err != nil {
		logger.Errorf("get path dir error: %v", err)
		return nil
	}
	return dirs
}

func (f *FileSysUtils) GetHashWithFilePath(filePath string) string {
	hash, err := fileSystem.GetHashWithFilePath(filePath)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return ""
	}
	return hash
}

func (f *FileSysUtils) GetHashsWithFilePaths(filePaths []string) []string {
	hashes, err := fileSystem.GetHashsWithFilePaths(filePaths)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return nil
	}
	return hashes
}

func (f *FileSysUtils) IntoFile(fileName string, data string) error {
	// 创建一个文件，然后写入数据
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

func (f *FileSysUtils) AddToFile(fileName string, data string) error {
	// 创建一个文件，然后写入数据
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

func (f *FileSysUtils) StrSpliceIntoFile(fileName string, strSlice []string) {
	// 创建一个文件，然后写入数据
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
