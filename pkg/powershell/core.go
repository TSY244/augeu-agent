package powershell

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func RunExec(cmd string) (string, error) {
	execRet := exec.Command("powershell", "-Command", cmd)
	output, err := execRet.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), err
}

// GetScheduledTaskCommands 使用 PowerShell 获取任务命令
func GetScheduledTaskCommands() ([]string, error) {
	cmd := exec.Command("powershell", "-Command", "Get-ScheduledTask | ForEach-Object { $_.Actions.Execute + ' ' + $_.Actions.Arguments }")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("执行 PowerShell 失败: %w", err)
	}

	// 解析输出
	lines := strings.Split(string(output), "\n")
	ret := make([]string, 0)
	for _, line := range lines {
		trimed := strings.TrimSpace(line)
		if trimed != "" {
			//fmt.Println(line)
			ret = append(ret, trimed)
		}
	}
	//fmt.Println(results)
	return ret, nil
}

func GetBitsAdminInfo() ([]BitsTask, error) {
	powershellCmd := `
	Import-Module BitsTransfer
	Get-BitsTransfer | Select-Object JobId, DisplayName, TransferType, JobState | Format-Table -AutoSize
	`

	// 创建一个 exec.Command 对象
	cmd := exec.Command("powershell.exe", "-Command", powershellCmd)

	// 捕获标准输出和标准错误
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatalf("执行 PowerShell 命令失败: %s\n错误信息: %s", err, stderr.String())
	}

	// 解析 PowerShell 输出
	tasks, err := parseBitsTransferOutput(stdout.String())
	if err != nil {
		log.Fatalf("解析 BITS 任务输出失败: %s", err)
	}
	returnTasks := make([]BitsTask, 0)
	targetLen := len("e701bb60-4b04-49a7-8086-51f6b01b723d")
	for _, task := range tasks {
		if len(task.JobId) != targetLen {
			continue
		}
		returnTasks = append(returnTasks, task)
	}
	return returnTasks, nil
}

func parseBitsTransferOutput(output string) ([]BitsTask, error) {
	var tasks []BitsTask

	// 使用 bufio.Scanner 按行读取输出
	scanner := bufio.NewScanner(strings.NewReader(output))
	isHeader := true // 标记是否是表头

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和表头
		if line == "" || isHeader {
			isHeader = false
			continue
		}

		// 按空格分割字段（假设字段之间有固定空格）
		fields := splitByWhitespace(line)
		if len(fields) < 4 {
			continue // 忽略格式不正确的行
		}

		// 构造 BitsTask 实例
		task := BitsTask{
			JobId:        fields[0],
			DisplayName:  fields[1],
			TransferType: fields[2],
			JobState:     fields[3],
		}

		// 添加到任务列表
		tasks = append(tasks, task)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取输出时出错: %w", err)
	}

	return tasks, nil
}

// 按空格分割字符串，忽略多余空格
func splitByWhitespace(input string) []string {
	var fields []string
	var currentField strings.Builder

	for _, char := range input {
		if char == ' ' || char == '\t' {
			if currentField.Len() > 0 {
				fields = append(fields, currentField.String())
				currentField.Reset()
			}
		} else {
			currentField.WriteRune(char)
		}
	}

	// 添加最后一个字段
	if currentField.Len() > 0 {
		fields = append(fields, currentField.String())
	}

	return fields
}

func T() {
	// 定义要执行的 PowerShell 命令
	powershellCmd := `
	Import-Module BitsTransfer
	Get-BitsTransfer | Select-Object JobId, DisplayName, TransferType, JobState | Format-Table -AutoSize
	`

	// 创建一个 exec.Command 对象
	cmd := exec.Command("powershell.exe", "-Command", powershellCmd)

	// 捕获标准输出和标准错误
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatalf("执行 PowerShell 命令失败: %s\n错误信息: %s", err, stderr.String())
	}

	// 解析 PowerShell 输出
	tasks, err := parseBitsTransferOutput(stdout.String())
	if err != nil {
		log.Fatalf("解析 BITS 任务输出失败: %s", err)
	}

	// 结构化输出任务信息
	fmt.Println("BITS 任务信息:")
	for _, task := range tasks {
		fmt.Printf("任务 ID: %s\n", task.JobId)
		fmt.Printf("显示名称: %s\n", task.DisplayName)
		fmt.Printf("传输类型: %s\n", task.TransferType)
		fmt.Printf("任务状态: %s\n", task.JobState)
		fmt.Println("-----------------------------")
	}
}
