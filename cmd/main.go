package main

import (
	"augeu-agent/internal/pkg/agent"
	"augeu-agent/internal/pkg/param"
)

func main() {
	param.Init()
	c := agent.Config{
		ConfigPath: param.BaseConfig.ConfigPath,
		Mode:       param.BaseConfig.Mode,
	}
	augeu := agent.NewAgent(&c)
	augeu.Run()
}
