package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/powershell"
)

type Schedule struct {
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (s *Schedule) GetScheduledTaskCommands() []string {
	ret, err := powershell.GetScheduledTaskCommands()
	if err != nil {
		logger.Errorf("获取计划任务命令失败: %v", err)
		return nil
	} else {
		return ret
	}
}
