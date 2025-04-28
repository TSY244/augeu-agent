package engUtils

import (
	"augeu-agent/internal/utils/consts"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/registration"
	"golang.org/x/sys/windows/registry"
	"regexp"
)

type Reg struct {
}

// NewReg 创建注册表操作对象
//
// return:
//   - *Reg 注册表操作实例
func NewReg() *Reg {
	return &Reg{}
}

// GetPathSubKeys 获取路径下的子项
//
// params:
//   - path 注册表的路径
//
// return:
//   - []string 子项
//
// notice:
//  1. 发生错误时返回nil
func (r *Reg) GetPathSubKeys(path string) []string {
	names, err := registration.GetPathSubKeys(path)
	if err != nil {
		logger.Errorf("get path sub keys error: %v", err)
		return nil
	}
	return names
}

// GetRegPathValueNames 获取注册表路径下的所有值名称
//
// params:
//   - path 注册表路径
//
// return:
//   - []string 值名称列表
//
// notice:
//  1. 发生错误时返回nil
func (r *Reg) GetRegPathValueNames(path string) []string {
	names, err := registration.GetRegPathValueNames(path)
	if err != nil {
		logger.Errorf("get reg path value names error: %v", err)
		return nil
	}
	return names
}

// GetRegPathValue 获取注册表路径指定值的字符串值
//
// params:
//   - path 注册表路径
//   - name 值名称
//
// return:
//   - string 值内容，错误时返回consts.ErrKey
//
// notice:
//  1. 非字符串类型值将返回错误
//  2. 错误时返回预定义错误标识consts.ErrKey
func (r *Reg) GetRegPathValue(path string, name string) string {
	value, err := registration.GetRegPathStringValue(path, name)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return value
}

// RegNameType 获取注册表键中指定值的类型
//
// params:
//   - reg 注册表键对象
//   - name 值名称
//
// return:
//   - uint32 注册表值类型代码（如registry.REG_SZ等）
//
// notice:
//  1. 错误时返回0
func (r *Reg) RegNameType(reg registry.Key, name string) uint32 {
	ret, err := registration.RegNameType(reg, name)
	if err != nil {
		logger.Errorf("get reg name type error: %v", err)
		return 0
	}
	return ret
}

// IsPathWithSubKey 检查注册表路径是否存在指定子键
//
// params:
//   - path 父路径
//   - subKey 子键名称
//
// return:
//   - bool 是否存在
func (r *Reg) IsPathWithSubKey(path string, subKey string) bool {
	return registration.IsPathWithSubKey(path, subKey)
}

// IsPathWithName 检查注册表路径是否存在指定值名称
//
// params:
//   - path 注册表路径
//   - name 值名称
//
// return:
//   - bool 是否存在
func (r *Reg) IsPathWithName(path string, name string) bool {
	return registration.IsPathWithName(path, name)
}

// GetPathFromCmd 从命令字符串中提取路径
//
// params:
//   - cmd 命令字符串（如"\"C:\Program Files\app.exe\" arg1\"）
//
// return:
//   - string 提取的路径（如"C:\Program Files\app.exe"）
//
// notice:
//  1. 支持带引号和不带引号的路径格式
//  2. 未匹配时返回空字符串
func (r *Reg) GetPathFromCmd(cmd string) string {
	re := regexp.MustCompile(`^"([^"]+)"|^([^ "]+)`)
	matches := re.FindStringSubmatch(cmd)
	if len(matches) == 0 {
		return ""
	}
	if matches[1] != "" {
		return matches[1]
	}
	return matches[2]
}

// IsHavePath 检查注册表路径是否存在
//
// params:
//   - path 注册表路径
//
// return:
//   - bool 路径是否存在
func (r *Reg) IsHavePath(path string) bool {
	return registration.IsHavePath(path)
}

// GetDefaultRegPathValue 获取注册表路径的默认值
//
// params:
//   - path 注册表路径
//
// return:
//   - string 默认值内容，错误时返回consts.ErrKey
//
// notice:
//  1. 错误时返回预定义错误标识consts.ErrKey
func (r *Reg) GetDefaultRegPathValue(path string) string {
	ret, err := registration.GetDefaultRegPathValue(path)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return ret
}
