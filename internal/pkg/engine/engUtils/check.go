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

func NewCheck() *Check {
	return &Check{}
}

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
