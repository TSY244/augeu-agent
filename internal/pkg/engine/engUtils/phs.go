package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/powershell"
	"fmt"
)

type Phs struct {
}

func NewPhs() *Phs {
	return &Phs{}
}

func (p *Phs) GetScheduledTaskCommands() []string {
	ret, err := powershell.GetBitsAdminInfo()
	if err != nil {
		logger.Errorf("获取计划任务命令失败: %v", err)
		return nil
	} else {
		var retStr []string
		for _, v := range ret {
			strValue := fmt.Sprintf("id: %s 任务名: %s 传输类型: %s 任务状态: %s ",
				v.JobId, v.DisplayName, v.TransferType, v.JobState)
			retStr = append(retStr, strValue)
		}
		return retStr
	}
}
