package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/powershell"
	"fmt"
)

type Phs struct {
}

// NewPhs 创建 PowerShell 工具实例
//
// return:
//   - *Phs PowerShell 工具实例
func NewPhs() *Phs {
	return &Phs{}
}

// GetScheduledTaskCommands 获取计划任务的命令信息
//
// return:
//   - []string 包含计划任务命令信息的字符串列表，错误时返回nil
//
// notice:
//  1. 使用 powershell.GetBitsAdminInfo 获取计划任务信息
//  2. 如果获取信息失败，记录错误日志并返回nil
//  3. 每个计划任务的信息格式为："id: <JobId> 任务名: <DisplayName> 传输类型: <TransferType> 任务状态: <JobState>"
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
