package agent

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/registration"
	"golang.org/x/sys/windows/registry"
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
	//return registration.GetRegPathValue(path, name)
	value, err := registration.GetRegPathValue(path, name)
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
