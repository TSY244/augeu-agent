package registration

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"strings"
)

var (
	WindowsId = "" // 提高效率，只获取一次
)

// GetWindowsGuid 通过注册表获取windows 机器的guid
//
// 注意：
//   - 通过注册表获取的guid 只有当用户重新安装了系统，或者是修改了windows 的安装位置才会发生变化
func GetWindowsGuid() (string, error) {
	if WindowsId != "" {
		return WindowsId, nil
	}
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, GuidKeyPath, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()
	value, _, err := key.GetStringValue(GuidKeyName)
	if err != nil {
		return "", err
	}
	WindowsId = value
	return value, nil
}

// 使用wmi 进行替代
//// GetWindowsUserNamesByRegistry 获取windows 所有的用户名称
//func GetWindowsUserNamesByRegistry() ([]string, error) {
//	key, err := registry.OpenKey(
//		registry.LOCAL_MACHINE,
//		UserNamesPath,
//		registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE,
//	)
//	if err != nil {
//		return nil, fmt.Errorf("需要以管理员权限运行程序。错误: %v", err)
//	}
//	defer key.Close()
//
//	names, err := key.ReadValueNames(0)
//	if err != nil {
//		return nil, fmt.Errorf("读取失败: %v", err)
//	}
//	return names, nil
//}

// 获取一个path 下的所有的子key
func GetSubKeys(key registry.Key, path string) ([]string, error) {
	temp, err := registry.OpenKey(
		key,
		path,
		registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE,
	)
	if err != nil {
		return nil, err
	}
	defer temp.Close()
	names, err := temp.ReadSubKeyNames(0)
	if err != nil {
		return nil, err
	}
	return names, nil
}

func GetDebuggerValue(fileName string) (string, error) {
	path := IFEOPath + "\\" + fileName
	reg, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer reg.Close()
	value, _, err := reg.GetStringValue(IFEODebuggerKeyName)
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetPathSubKeys 获取一个path 下的所有的子key
func GetPathSubKeys(path string) ([]string, error) {
	root, otherPath, err := RegSplit(path)
	if err != nil {
		return nil, err
	}
	names, err := GetSubKeys(root, otherPath)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// GetRegPathValueNames 获取一个path 下的所有的value name
func GetRegPathValueNames(path string) ([]string, error) {
	root, otherPath, err := RegSplit(path)
	if err != nil {
		return nil, err
	}
	reg, err := registry.OpenKey(root, otherPath, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer reg.Close()
	names, err := reg.ReadValueNames(0)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// GetRegPathValue 获取一个path 对应的name 的value
func GetRegPathValue(path string, name string) (string, error) {
	root, otherPath, err := RegSplit(path)
	if err != nil {
		return "", err
	}
	reg, err := registry.OpenKey(root, otherPath, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer reg.Close()
	valueType, err := RegNameType(reg, name)
	if err != nil {
		return "", err
	}
	value, err := getRegPathValueBase(reg, name, valueType)
	if err != nil {
		return "", err
	}
	return value, nil
}

// getRegPathValueBase 一个通用的获取value 的方法
func getRegPathValueBase(reg registry.Key, name string, t uint32) (string, error) {
	switch t {
	case REG_SZ:
		value, _, err := reg.GetStringValue(name)
		if err != nil {
			return "", err
		}
		return value, nil
	case REG_DWORD:
		value, _, err := reg.GetIntegerValue(name)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", value), nil
	case REG_QWORD:
		value, _, err := reg.GetIntegerValue(name)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", value), nil
	case REG_BINARY:
		value, _, err := reg.GetBinaryValue(name)
		if err != nil {
			return "", err
		}
		return string(value), nil
	case REG_MULTI_SZ:
		value, _, err := reg.GetStringsValue(name)
		if err != nil {
			return "", err
		}
		return strings.Join(value, ","), nil
	case REG_EXPAND_SZ:
		value, _, err := reg.GetStringValue(name)
		if err != nil {
			return "", err
		}
		return value, nil
	default:
		return "", fmt.Errorf("unsupported type: %d", t)
	}
}

// RegNameType 获取一个name 的类型
func RegNameType(reg registry.Key, name string) (uint32, error) {
	_, t, err := reg.GetValue(name, nil)
	if err != nil {
		return 0, err
	}
	return t, nil
}

// IsPathWithName 判断一个reg path 是否包含对应的一个 name
func IsPathWithName(path string, name string) bool {
	names, err := GetRegPathValueNames(path)
	if err != nil {
		return false
	}
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}

// IsPathWithSubKey 判断一个reg path 是否包含对应的一个 sub key
func IsPathWithSubKey(path string, subKey string) bool {
	subKeys, err := GetPathSubKeys(path)
	if err != nil {
		return false
	}
	for _, n := range subKeys {
		if n == subKey {
			return true
		}
	}
	return false
}
