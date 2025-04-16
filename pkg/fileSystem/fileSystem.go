package fileSystem

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// LsFile 获取路径下的文件列表， 不包含文件夹
func LsFile(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

// LsDir 获取路径下的文件夹列表， 不包含文件
func LsDir(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var dirNames []string
	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
		}
	}
	return dirNames, nil
}

// GetHashWithFilePath  不支持文件夹
func GetHashWithFilePath(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// FromLinkToPath 获取链接指向的路径
func FromLinkToPath(linkPath string) (string, error) {
	file, err := os.Open(linkPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	return fileInfo.Mode().String(), nil
}
