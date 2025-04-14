package windowsWmi

type QueryString struct {
	Value string
}

type QueryStrings struct {
	Values []string
}

type QueryFunc func(query string) ([]string, error)

type QueryKey string

type Win32_QuickFixEngineering struct {
	Description string
	HotFixID    string
	InstalledBy string
	InstalledOn string
}

// 确保结构体字段与WMI属性映射
type Win32_ComputerSystemProduct struct {
	UUID string `wmi:"UUID"`
}

type Win32_OperatingSystem struct {
	Caption string `wmi:"Caption"`
}

type Win32_OperatingSystemVersion struct {
	Version string `wmi:"Version"`
}

type Win32_UserAccount struct {
	Name         string `wmi:"Name"`
	Description  string `wmi:"Description"`
	LocalAccount bool   `wmi:"LocalAccount"`
	SID          string `wmi:"SID"`
}

type win32_ScheduledTask struct {
	TaskName    string `wmi:"TaskName"`
	Author      string `wmi:"Author"`
	State       int    `wmi:"State"`
	Description string `wmi:"Description"`
	TaskPath    string `wmi:"TaskPath"`
	URI         string `wmi:"URI"`
}

type Win32_ScheduledTask struct {
	TaskName    string
	Author      string
	State       string
	Description string
	TaskPath    string
	URI         string
}

type Win32_Service struct {
	Name        string `wmi:"Name"`
	DisplayName string `wmi:"DisplayName"`
	State       string `wmi:"State"`
	StartMode   string `wmi:"StartMode"`
	StartName   string `wmi:"StartName"`
}
