package engUtils

import (
	"augeu-agent/pkg/windowsWmi"
	"regexp"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// GetServiceName 获取服务名称
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
