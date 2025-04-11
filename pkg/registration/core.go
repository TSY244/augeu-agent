package registration

import (
	"golang.org/x/sys/windows/registry"
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
