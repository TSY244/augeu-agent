package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/powershell"
)

type Schedule struct {
}

// NewSchedule 创建计划任务工具实例
//
// return:
//   - *Schedule 计划任务工具实例
func NewSchedule() *Schedule {
	return &Schedule{}
}

// GetScheduledTaskCommands 获取计划任务的命令列表
//
// return:
//   - []string 包含计划任务命令的字符串列表，错误时返回nil
//
// notice:
//  1. 使用 powershell.GetScheduledTaskCommands 获取计划任务命令
//  2. 如果获取命令失败，记录错误日志并返回nil
func (s *Schedule) GetScheduledTaskCommands() []string {
	ret, err := powershell.GetScheduledTaskCommands()
	if err != nil {
		logger.Errorf("获取计划任务命令失败: %v", err)
		return nil
	} else {
		return ret
	}
}
