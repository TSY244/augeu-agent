package engUtils

import (
	"augeu-agent/internal/pkg/weibu"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/util/convert"
	"fmt"
	"strings"
)

type Check struct {
}

// NewCheck 创建检查工具实例
//
// return:
//   - *Check 检查工具实例
func NewCheck() *Check {
	return &Check{}
}

// CheckHash 通过微步检查哈希值的风险率是否低于指定阈值
//
// params:
//   - hash 哈希值（如文件的MD5、SHA256等）
//   - a 配置对象，包含多引擎检测相关配置
//   - proxy 代理地址（可选，用于网络请求）
//   - rate 风险率阈值（如0.5表示50%）
//
// return:
//   - bool 是否通过检查（风险率 <= rate为true，否则为false）
//
// notice:
//  1. 如果获取多引擎检测结果失败，返回false
//  2. 如果返回结果格式不正确（如不包含"/"或分割后长度不为2），返回false
//  3. 如果字符串转整数失败，返回false
//  4. 风险率计算公式：intRate1 / intRate2，其中intRate1和intRate2为分割后的两个整数
func (c *Check) CheckHash(hash string, a *weibu.Config, proxy string, rate float64) bool {
	ret, err := weibu.GetMultiEngines(hash, a, proxy)
	if err != nil {
		logger.Errorf("CheckHash error is %v", err)
		return false
	}
	if ret == "" {
		return false
	}
	if !strings.Contains(ret, "/") {
		logger.Errorf("CheckHash error is %v", fmt.Errorf("ret is %s", ret))
		return false
	}
	rates := strings.Split(ret, "/")
	if len(rates) != 2 {
		logger.Errorf("CheckHash error is %v", fmt.Errorf("ret is %s", ret))
		return false
	}
	intRate1, err := convert.Str2Int(rates[0])
	if err != nil {
		logger.Errorf("CheckHash error is %v", err)
		return false
	}
	intRate2, err := convert.Str2Int(rates[1])
	if err != nil {
		logger.Errorf("CheckHash error is %v", err)
		return false
	}
	return float64(intRate1)/float64(intRate2) <= rate
}
