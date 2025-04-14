package main

import (
	"augeu-agent/internal/pkg/agent"
	"augeu-agent/internal/pkg/param"
)

func main() {
	param.Init()

	c := agent.Config{}
	augeu := agent.NewAgent(&c)
	augeu.Run()
}
