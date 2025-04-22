package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/registration"
	"golang.org/x/sys/windows/registry"
	"regexp"
)

type Reg struct {
}

func NewReg() *Reg {
	return &Reg{}
}

func (r *Reg) GetPathSubKeys(path string) []string {
	names, err := registration.GetPathSubKeys(path)
	if err != nil {
		logger.Errorf("get path sub keys error: %v", err)
		return nil
	}
	return names
}

func (r *Reg) GetRegPathValueNames(path string) []string {
	//return registration.GetRegPathValueNames(path)
	names, err := registration.GetRegPathValueNames(path)
	if err != nil {
		logger.Errorf("get reg path value names error: %v", err)
		return nil
	}
	return names
}

func (r *Reg) GetRegPathValue(path string, name string) string {
	//return registration.GetRegPathStringValue(path, name)
	value, err := registration.GetRegPathStringValue(path, name)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return ""
	}
	return value
}

func (r *Reg) RegNameType(reg registry.Key, name string) uint32 {
	//return registration.RegNameType(reg, name)
	ret, err := registration.RegNameType(reg, name)
	if err != nil {
		logger.Errorf("get reg name type error: %v", err)
		return 0
	}
	return ret
}

func (r *Reg) IsPathWithSubKey(path string, subKey string) bool {
	return registration.IsPathWithSubKey(path, subKey)
}
func (r *Reg) IsPathWithName(path string, name string) bool {
	return registration.IsPathWithName(path, name)
}

func (r *Reg) GetPathFromCmd(cmd string) string {
	re := regexp.MustCompile(`^"([^"]+)"|^([^ "]+)`) // 匹配带引号和不带引号的路径
	matches := re.FindStringSubmatch(cmd)
	if len(matches) == 0 {
		return ""
	}

	var path string
	if matches[1] != "" { // 带引号的路径
		path = matches[1]
	} else { // 不带引号的路径
		path = matches[2]
	}
	return path
}

func (r *Reg) IsHavePath(path string) bool {
	return registration.IsHavePath(path)
}

func (r *Reg) GetDefaultRegPathValue(path string) string {
	ret, err := registration.GetDefaultRegPathValue(path)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return ""
	}
	return ret
}
