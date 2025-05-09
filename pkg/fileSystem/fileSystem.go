package fileSystem

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

// LsFile 获取路径下的文件列表， 不包含文件夹
func LsFile(path string) ([]string, error) {
	path = replaceEnvPath(path)
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
	path = replaceEnvPath(path)
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
	filePath = replaceEnvPath(filePath)
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

func GetHashsWithFilePaths(filePaths []string) ([]string, error) {

	var hashs []string
	for _, filePath := range filePaths {
		filePath = replaceEnvPath(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		hash := md5.New()
		_, err = io.Copy(hash, file)
		if err != nil {
			return nil, err
		}
		hashs = append(hashs, hex.EncodeToString(hash.Sum(nil)))
		file.Close()
	}
	return hashs, nil
}

// FromLinkToPath 获取链接指向的路径
//
// 注意：
//  1. windows 不工作
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

func replaceEnvPath(path string) string {
	// %appdata%\Microsoft\Windows\Start Menu\Programs\Startup
	if strings.Contains(path, "%") {
		splitedPaths := strings.Split(path, "%")
		if len(splitedPaths) != 3 {
			return path
		}
		env := splitedPaths[1]
		envPath := os.Getenv(env)
		return envPath + splitedPaths[2]
	}
	return path
}
