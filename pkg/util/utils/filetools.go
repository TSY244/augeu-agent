package utils

import (
	"os"
	"strings"
)

const (
	// Delimiter is the separator used to separate different screenshots
	Delimiter = "\n\n---SEPARATOR---\n\n"
	CacheSize = 1024 * 8
)

func ReadFile(path string) ([]byte, error) {
	// check file is exist
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteFile(path string, data []byte, perm os.FileMode) error {
	err := os.WriteFile(path, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CheckFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetFileNameFormPath(path string) string {
	if strings.Contains(path, "\\") {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path[strings.LastIndex(path, "/")+1:]
}
