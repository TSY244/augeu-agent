package engUtils

import (
	"augeu-agent/pkg/windowsWmi"
	"regexp"
)

type Service struct {
}

// NewService 创建服务工具实例
//
// return:
//   - *Service 服务工具实例
func NewService() *Service {
	return &Service{}
}

// GetServiceName 获取所有服务的名称
//
// return:
//   - []string 包含服务名称的字符串列表，错误时返回空列表
//
// notice:
//  1. 使用 windowsWmi.QueryServicesDetail 获取服务详细信息
//  2. 如果查询失败，返回空列表
func (s *Service) GetServiceName() []string {
	names := make([]string, 0)
	ret, err := windowsWmi.QueryServicesDetail()
	if err != nil {
		return names
	}
	for _, v := range ret {
		names = append(names, v.Name)
	}
	return names
}

// GetServiceCmd 根据服务名称获取其启动命令
//
// params:
//   - name 服务名称
//
// return:
//   - string 服务的启动命令路径，未找到或错误时返回空字符串
//
// notice:
//  1. 使用 windowsWmi.QueryServicesDetail 获取服务详细信息
//  2. 如果查询失败或未找到对应服务，返回空字符串
func (s *Service) GetServiceCmd(name string) string {
	ret, err := windowsWmi.QueryServicesDetail()
	if err != nil {
		return ""
	}
	for _, v := range ret {
		if v.Name == name {
			return v.PathName
		}
	}
	return ""
}

// GetRunningServiceCmd 获取所有正在运行的服务的启动命令
//
// return:
//   - []string 包含正在运行服务的启动命令路径的字符串列表，错误时返回空列表
//
// notice:
//  1. 使用 windowsWmi.QueryServicesDetail 获取服务详细信息
//  2. 仅返回 StartMode 为 "Running" 的服务的启动命令
//  3. 如果查询失败，返回空列表
func (s *Service) GetRunningServiceCmd() []string {
	cmds := make([]string, 0)
	ret, err := windowsWmi.QueryServicesDetail()
	if err != nil {
		return cmds
	}
	for _, v := range ret {
		if v.StartMode == "Running" {
			cmds = append(cmds, v.PathName)
		}
	}
	return cmds
}

// GetServiceImagePath 获取所有服务的镜像路径
//
// return:
//   - []string 包含服务镜像路径的字符串列表，错误时返回空列表
//
// notice:
//  1. 使用 windowsWmi.QueryServicesDetail 获取服务详细信息
//  2. 调用 GetPathFromCmd 提取 PathName 中的路径部分
//  3. 如果查询失败，返回空列表
func (s *Service) GetServiceImagePath() []string {
	imgPath := make([]string, 0)
	ret, err := windowsWmi.QueryServicesDetail()
	if err != nil {
		return imgPath
	}
	for _, v := range ret {
		imgPath = append(imgPath, GetPathFromCmd(v.PathName))
	}
	return imgPath
}

// GetPathFromCmd 从命令字符串中提取路径
//
// params:
//   - cmd 命令字符串（如 "\"C:\\Program Files\\app.exe\" arg1"）
//
// return:
//   - string 提取的路径（如 "C:\\Program Files\\app.exe"），未匹配时返回空字符串
//
// notice:
//  1. 支持带引号和不带引号的路径格式
//  2. 使用正则表达式匹配路径部分
func GetPathFromCmd(cmd string) string {
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
