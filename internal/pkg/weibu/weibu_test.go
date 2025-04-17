package weibu

import (
	"augeu-agent/internal/pkg/agent"
	"augeu-agent/internal/pkg/param"
	"augeu-agent/internal/utils/consts"
	"testing"
)

func TestGetFilesReport(t *testing.T) {

	agentConf := param.Config{
		Mode:        consts.BasicMode,
		WeiBuApiKey: "",
	}

	augeu := agent.NewAgent(&agentConf)

	targets := []string{
		"2c06f3f9a3bbcfe80f266816389a6c17",
		"d47ab4565522701a772af1af2863d85d",
		"ebc6bfd15c7bcc0b9538878f05e80576",
	}
	ret, err := GetFilesReport(targets, augeu.GetWeiBuConf())
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
