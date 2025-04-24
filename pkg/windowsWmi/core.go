package windowsWmi

import (
	"errors"
	"fmt"
	"github.com/StackExchange/wmi"
	"log"
)

// QueryUuid 查询系统UUID
func QueryUuid() (string, error) {
	var dst []Win32_ComputerSystemProduct
	if err := wmi.Query(QueryUuidKey, &dst); err != nil {
		return "", fmt.Errorf("WMI query failed: %w", err)
	}
	if len(dst) == 0 {
		return "", errors.New("empty result set")
	}
	return dst[0].UUID, nil
}

// QueryOsName 查询系统名称
func QueryOsName() (string, error) {
	var dst []Win32_OperatingSystem
	if err := wmi.Query(QueryOsNameKey, &dst); err != nil {
		return "", fmt.Errorf("WMI query failed: %w", err)
	}
	if len(dst) == 0 {
		return "", errors.New("empty result set")
	}
	return dst[0].Caption, nil
}

// QueryOsVersion 查询系统版本
func QueryOsVersion() (string, error) {
	var dst []Win32_OperatingSystemVersion
	if err := wmi.Query(QueryOsVersionKey, &dst); err != nil {
		return "", fmt.Errorf("WMI query failed: %w", err)
	}
	if len(dst) == 0 {
		return "", errors.New("empty result set")
	}
	return dst[0].Version, nil
}

// 统一使用基础查询函数
func baseSlice[T any](resultSet *[]T, queryKey string) error {
	if resultSet == nil {
		return errors.New("nil pointer passed for result storage")
	}
	if err := wmi.Query(queryKey, resultSet); err != nil {
		return fmt.Errorf("WMI query failed: %w", err)
	}
	if len(*resultSet) == 0 {
		return errors.New("empty result set from WMI query")
	}
	return nil
}

// QueryHotFix 查询系统补丁
func QueryHotFix() ([]Win32_QuickFixEngineering, error) {
	var dst []Win32_QuickFixEngineering
	if err := baseSlice(&dst, QueryHotFixKey); err != nil {
		return nil, err
	}
	return dst, nil
}

func QueryUsers() ([]Win32_UserAccount, error) {
	var dst []Win32_UserAccount
	if err := baseSlice(&dst, QueryUsersKey); err != nil {
		return nil, err
	}
	return dst, nil
}

func QueryScheduledTasks() ([]Win32_ScheduledTask, error) {
	namespace := "Root\\Microsoft\\Windows\\TaskScheduler"

	// 准备接收查询结果
	var tasks []win32_ScheduledTask

	// 执行 WMI 查询
	query := wmi.CreateQuery(&tasks, "", "MSFT_ScheduledTask")
	err := wmi.QueryNamespace(query, &tasks, namespace)
	if err != nil {
		log.Fatalf("WMI 查询失败: %v (可能需要管理员权限)", err)
	}

	stateMap := map[int]string{
		1: "禁用",
		3: "就绪",
		4: "运行中",
	}

	var result []Win32_ScheduledTask

	for _, task := range tasks {
		state, ok := stateMap[task.State]
		if !ok {
			state = "未知"
		}
		result = append(result, Win32_ScheduledTask{
			TaskName:    task.TaskName,
			Author:      task.Author,
			State:       state,
			Description: task.Description,
			TaskPath:    task.TaskPath,
			URI:         task.URI,
		})
	}
	return result, nil
}

func QueryServices() ([]Win32_Service, error) {
	var dst []Win32_Service
	if err := baseSlice(&dst, QueryServiceKey); err != nil {
		return nil, err
	}
	return dst, nil
}

// QueryServicesDetail 查询所有Windows服务
func QueryServicesDetail() ([]Win32_ServiceDetail, error) {
	var dst []Win32_ServiceDetail
	if err := baseSlice(&dst, QueryServiceDetailKey); err != nil {
		return nil, err
	}
	return dst, nil
}
