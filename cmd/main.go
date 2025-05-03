package main

import (
	"augeu-agent/internal/pkg/agent"
	"augeu-agent/internal/pkg/param"
)

func main() {
	param.Init()
	augeu := agent.NewAgent(&param.BaseConfig)
	augeu.SetEnv()
	augeu.Run()
}
