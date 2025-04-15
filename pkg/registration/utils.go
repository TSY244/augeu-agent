package registration

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"strings"
)

func RegSplit(path string) (registry.Key, string, error) {
	index := strings.Index(path, "\\")
	if index == -1 {
		return registry.LOCAL_MACHINE, "", fmt.Errorf("path not found: %s", path)
	}
	rootPath := path[:index]
	otherPath := path[index+1:]
	root, ok := RootPathMap[rootPath]
	if !ok {
		return registry.LOCAL_MACHINE, "", fmt.Errorf("root path not found: %s", rootPath)
	}

	return root, otherPath, nil
}
