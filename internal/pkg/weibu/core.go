package weibu

import (
	"augeu-agent/internal/utils/consts"
	"augeu-agent/pkg/augeuHttp"
	"augeu-agent/pkg/util/utils"
	"encoding/json"
	"fmt"
)

func GetFileReport(target string, a *Config, proxy string) (string, error) {
	url := fmt.Sprintf("https://api.threatbook.cn/v3/file/report?apikey=%s&sandbox_type=win10_1903_enx64_office2016&resource=%s&query_fields=summary&query_fields=multiengines", a.Conf.WeiBuApiKey, target)
	if a.Conf.Mode != consts.RemoteModeApi {
		ret, err := augeuHttp.GetRequest(url, a.Header, proxy)
		if err != nil {
			return "", err
		}
		// json data -> response
		wbResp := weiBuResponse{}
		err = json.Unmarshal([]byte(ret), &wbResp)
		if err != nil {
			return "", err
		}
		strData := convert2WeiBuDataChinese(&wbResp)
		//logger.Debugf("%v", strData)
		return strData, nil
	}
	return "", nil
}

// GetFilesReport 使用代理获取多个文件的报告
//
// 参数：
//   - target: 目标文件列表
//   - agent: 代理对象
//   - proxys: 代理列表，是随机选择的
func GetFilesReport(targets []string, c *Config, proxys ...string) ([]string, error) {
	nouse := false
	if len(proxys) == 0 {
		nouse = true
	}
	resp := make([]string, 0)
	for _, target := range targets {
		proxy := ""
		if !nouse {
			proxy = proxys[utils.GetRandom(0, len(proxys))]
		}
		ret, err := GetFileReport(target, c, proxy)
		if err != nil {
			return nil, err
		}
		resp = append(resp, ret)
	}
	return resp, nil
}

func GetMultiEngines(target string, a *Config, proxy string) (string, error) {
	url := fmt.Sprintf("https://api.threatbook.cn/v3/file/report?apikey=%s&sandbox_type=win10_1903_enx64_office2016&resource=%s&query_fields=summary&query_fields=multiengines", a.Conf.WeiBuApiKey, target)
	if a.Conf.Mode != consts.RemoteModeApi {
		ret, err := augeuHttp.GetRequest(url, a.Header, proxy)
		if err != nil {
			return "", err
		}
		// json data -> response
		wbResp := weiBuResponse{}
		err = json.Unmarshal([]byte(ret), &wbResp)
		if err != nil {
			return "", err
		}
		return wbResp.Data.Summary.MultiEngines, nil
	}
	return "", nil
}
