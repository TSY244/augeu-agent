package fileSystem

import "testing"

func TestLsFile(t *testing.T) {
	filePath := "../../../"
	fileList, err := LsFile(filePath)
	if err != nil {
		t.Error(err)
	}
	for _, file := range fileList {
		t.Log(file)
	}
}

func TestLsDir(t *testing.T) {
	filePath := "../../../"
	fileList, err := LsDir(filePath)
	if err != nil {
		t.Error(err)
	}
	for _, file := range fileList {
		t.Log(file)
	}
}

func TestGetHashWithFilePath(t *testing.T) {
	filePath := "D:\\application\\Dism++\\Dism++x64.exe"
	fileSize, err := GetHashWithFilePath(filePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(fileSize)
}

func TestFromLinkToPath(t *testing.T) {
	linkPath := "C:\\Users\\12414\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\飞书.lnk"
	filePath, err := FromLinkToPath(linkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(filePath)
}
