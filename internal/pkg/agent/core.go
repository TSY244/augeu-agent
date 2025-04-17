package agent

import (
	"augeu-agent/internal/pkg/engine/engUtils"
	"augeu-agent/internal/pkg/param"
	"augeu-agent/pkg/color"
	"augeu-agent/pkg/logger"
	"fmt"
	engine2 "github.com/bilibili/gengine/engine"
	"os"
	"time"
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
			"println":      fmt.Println,
			"reg":          engUtils.NewReg(),
			"strUtils":     engUtils.NewStrUtils(),
			"fileSysUtils": engUtils.NewFileSys(),
			"printer":      engUtils.NewPrinter(),
			"base":         engUtils.NewBase(),
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
	startTime := time.Now().UnixNano()
	err, mapData := engPool.ExecuteConcurrent(nil)
	if err != nil {
		return err
	}
	endTime := time.Now().UnixNano()
	var falseRet []string
	info := fmt.Sprintf("-----------------------------------------------扫描完成 扫描耗时: %d ms-----------------------------------------------", (endTime-startTime)/1000000)
	color.Blue("%s\n", info)
	logger.Infof("返回结果:")
	for k, v := range mapData {
		logger.Infof("rule name: %s -> ret: %v", k, v)
		ret, ok := v.(bool)
		if ok && !ret {
			falseRet = append(falseRet, k)
		}
	}

	//logger.Infof("需要注意的rule:")
	color.Magenta("%s\n", "需要注意的rule:")
	for _, v := range falseRet {
		//logger.Warnf("rule name: %s", v)
		color.White("%s", "rule name:")
		color.HRed(" %s\n", v)
	}
	return nil
}

func (a *Agent) basicRun() {
	err := a.baseRun()
	if err != nil {
		logger.Errorf("basic run error: %v", err)
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
