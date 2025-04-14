package agent

import (
	"augeu-agent/internal/pkg/param"
	"augeu-agent/pkg/logger"
)

type Agent struct {
	Conf *Config
}

type Config struct {
	ConfigPath string
	Mode       string
}

func NewAgent(c *Config) *Agent {
	return &Agent{
		Conf: c,
	}
}

func (agent *Agent) Run() {
	switch agent.Conf.Mode {
	case BasicMode:
		logger.Infof("basic mode")
	case RemoteMode:
		logger.Infof("remote mode")
	case LocalMode:
		logger.Infof("local mode")
	default:
		logger.Errorf("unknown mode")
		param.Help()
	}
}
