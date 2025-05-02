package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/windowsApi"
	"fmt"
)

type ProcessUtils struct {
}

func NewProcessUtils() *ProcessUtils {
	return &ProcessUtils{}
}

func (p *ProcessUtils) SearchMem(target string) []string {
	var returnedString []string
	processes, err := windowsApi.GetAllProcess()
	if err != nil {
		logger.Errorf("GetAllProcess failed: %v", err)
	}
	for _, process := range processes {
		hProcess, err := windowsApi.OpenProcess(process.Pid)
		if err != nil {
			logger.Errorf("OpenProcess failed: %v", err)
		}

		ret, err := windowsApi.ScanProcessMemory(hProcess, "123&testaugust..redxcftgvyhbjnkmaskldjfhalks;jdf290318u7498uiojhsadnfjhbnaslkdjfh123&testaugust..redxcftgvyhbjnkmaskldjfhalks;jdf290318u7498uiojhsadnfjhbnaslkdjfh123&testaugust..redxcftgvyhbjnkmaskldjfhalks;jdf290318u7498uiojhsadnfjhbnaslkdjfh")
		if err != nil {
			logger.Errorf("SearchProcessMem failed: %v", err)
		}
		if ret {
			exePath, err := windowsApi.GetExePath(process.Pid)
			if err != nil {
				logger.Errorf("GetExePath failed: %v", err)
			}
			logger.Infof("Process Name: %s, Pid: %d, ExePath: %s", process.Name, process.Pid, exePath)
			returnedString = append(returnedString, fmt.Sprintf("Process Name: %s, Pid: %d, ExePath: %s Found %s", process.Name, process.Pid, exePath, target))
		}
	}
	return returnedString
}
