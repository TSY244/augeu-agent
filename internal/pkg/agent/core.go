package agent

import (
	"augeu-agent/internal/pkg/engine/engUtils"
	"augeu-agent/internal/pkg/param"
	"augeu-agent/pkg/logger"
	"fmt"
	engine2 "github.com/bilibili/gengine/engine"
	"os"
)

type Agent struct {
	Conf *Config
	//Eng      *engine.Engine
	Rule     string
	ApiOuter map[string]interface{}
}

type Config struct {
	ConfigPath string
	Mode       string
}

func NewAgent(c *Config) *Agent {
	return &Agent{
		Conf: c,
		//Eng:  engine.NewEngine(),
		ApiOuter: map[string]interface{}{
			"println":  fmt.Println,
			"reg":      NewReg(),
			"strUtils": engUtils.NewStrUtils(),
		},
	}
}

func (a *Agent) Run() {
	//	a.Rule = `
	//rule "name test" "i can"  salience 0
	//begin
	//	println("asdfasdf")
	//end
	//`

	switch a.Conf.Mode {
	case BasicMode:
		logger.Infof("basic mode")
		a.basicRun()
	case RemoteMode:
		logger.Infof("remote mode")
		a.remoteRun()
	case LocalMode:
		logger.Infof("local mode")
		a.localRun()
	default:
		logger.Errorf("unknown mode")
		param.Help()
	}
}

func (a *Agent) SetRule(rule string) {
	a.Rule = rule
}

func (a *Agent) baseRun() error {
	//a.AddObject()
	engPool, err := engine2.NewGenginePool(MinLen, MaxLen, Parallel, a.Rule, a.ApiOuter)
	//err := a.Eng.LoadRule(a.Rule)
	if err != nil {
		return err
	}
	err, mapData := engPool.ExecuteConcurrent(nil)
	if err != nil {
		return err
	}
	logger.Infof("返回结果: %v", mapData)
	return nil
}

func (a *Agent) basicRun() {
	err := a.baseRun()
	if err != nil {
		logger.Infof("basic run error: %v", err)
		os.Exit(1)
	}
}

func (a *Agent) remoteRun() {
	err := a.baseRun()
	if err != nil {
		logger.Errorf("basic run error: %v", err)
		os.Exit(1)
	}
}

func (a *Agent) localRun() {
	err := a.baseRun()
	if err != nil {
		logger.Infof("basic run error: %v", err)
		os.Exit(1)
	}
}
