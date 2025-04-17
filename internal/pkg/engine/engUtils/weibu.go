package engUtils

import (
	"augeu-agent/internal/pkg/weibu"
)

type WeiBuUtils struct {
}

func NewWeiBuUtils() *WeiBuUtils {
	return &WeiBuUtils{}
}

func (wb *WeiBuUtils) GetFileReport(target string, a *weibu.Config, proxy string) (string, error) {
	return weibu.GetFileReport(target, a, proxy)
}

func (wb *WeiBuUtils) GetFilesReport(targets []string, a *weibu.Config, proxy ...string) ([]string, error) {
	return weibu.GetFilesReport(targets, a, proxy...)
}
