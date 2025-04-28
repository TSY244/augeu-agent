package engUtils

import (
	"augeu-agent/internal/pkg/weibu"
)

type WeiBuUtils struct {
}

// NewWeiBuUtils 创建 WeiBu 工具实例
//
// return:
//   - *WeiBuUtils WeiBu 工具实例
func NewWeiBuUtils() *WeiBuUtils {
	return &WeiBuUtils{}
}

// GetFileReport 获取单个文件的报告
//
// params:
//   - target 目标文件路径或标识符
//   - a 配置对象，包含与 WeiBu 服务交互所需的参数
//   - proxy 代理地址（可选）
//
// return:
//   - string 文件报告的结果
//   - error 如果获取报告失败，返回错误信息；否则返回 nil
//
// notice:
//  1. 调用 weibu.GetFileReport 方法获取单个文件的报告
//  2. 如果发生错误，返回错误信息以便调用方处理
func (wb *WeiBuUtils) GetFileReport(target string, a *weibu.Config, proxy string) (string, error) {
	return weibu.GetFileReport(target, a, proxy)
}

// GetFilesReport 获取多个文件的报告
//
// params:
//   - targets 目标文件路径或标识符的切片
//   - a 配置对象，包含与 WeiBu 服务交互所需的参数
//   - proxy 可选的代理地址（支持多个代理地址）
//
// return:
//   - []string 文件报告结果的切片
//   - error 如果获取报告失败，返回错误信息；否则返回 nil
//
// notice:
//  1. 调用 weibu.GetFilesReport 方法获取多个文件的报告
//  2. 支持传入多个代理地址（通过可变参数 proxy 实现）
//  3. 如果发生错误，返回错误信息以便调用方处理
func (wb *WeiBuUtils) GetFilesReport(targets []string, a *weibu.Config, proxy ...string) ([]string, error) {
	return weibu.GetFilesReport(targets, a, proxy...)
}
