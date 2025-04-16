package engUtils

import (
	"augeu-agent/pkg/fileSystem"
	"augeu-agent/pkg/logger"
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
