package agent

import (
	"augeu-agent/internal/pkg/engine/engUtils"
	"augeu-agent/internal/pkg/param"
	"augeu-agent/internal/utils/consts"
	"augeu-agent/pkg/color"
	"augeu-agent/pkg/logger"
	"context"
	"fmt"
	engine2 "github.com/bilibili/gengine/engine"
	"os"
	"time"
)

type Agent struct {
	RootCtx context.Context
	Cancel  context.CancelFunc
	Conf    *param.Config
	//Eng      *engine.Engine
	Rule     string
	ApiOuter map[string]interface{}
	ClientId string
	Jwt      string
	Header   map[string]string
}

func NewAgent(c *param.Config) *Agent {
	checkConf(c)
	rootCtx, cancel := context.WithCancel(context.Background())

	//// websocket
	//ws, resp, err := websocket.DefaultDialer.Dial(c.WebsocketAddr, nil)
	//if err != nil {
	//	logger.Errorf("Failed to dial websocket: %v", err)
	//	cancel()
	//	return nil
	//}
	//if resp.StatusCode != 101 {
	//	logger.Errorf("Failed to dial websocket: %v", resp.Status)
	//	cancel()
	//	return nil
	//}
	//go func() {
	//	for {
	//		select {
	//		case <-rootCtx.Done():
	//			logger.Infof("websocket connection closed")
	//			return
	//		default:
	//		}
	//
	//		_, _, err := ws.ReadMessage()
	//		if err != nil {
	//			logger.Errorf("Lost Connection to server: %v", err)
	//			cancel()
	//			return
	//		}
	//
	//	}
	//}()
	agent := &Agent{
		Conf:    c,
		RootCtx: rootCtx,
		Cancel:  cancel,
		//Eng:  engine.NewEngine(),
	}
	agent.ApiOuter = map[string]interface{}{
		"println":      fmt.Println,
		"reg":          engUtils.NewReg(),
		"strUtils":     engUtils.NewStrUtils(),
		"fileSysUtils": engUtils.NewFileSys(),
		"printer":      engUtils.NewPrinter(),
		"base":         engUtils.NewBase(),
		"agent":        agent,
		"weibu":        engUtils.NewWeiBuUtils(),
	}
	return agent
}

func (a *Agent) Run() {
	switch a.Conf.Mode {
	case consts.BasicMode:
		a.SetRule(BasicRule)
		logger.Infof("basic mode")
		a.basicRun()
	case consts.RemoteMode:
		a.receiveClientId()
		logger.Infof("remote mode")
		a.remoteRun()
	case consts.LocalMode:
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
	engPool, err := engine2.NewGenginePool(consts.MinLen, consts.MaxLen, consts.Parallel, a.Rule, a.ApiOuter)
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

func checkConf(c *param.Config) {
	if c == nil {
		logger.Errorf("config is nil")
		param.Help()
		os.Exit(1)
	}
	if c.Mode == "" {
		logger.Errorf("mode is empty")
		param.Help()
		os.Exit(1)
	} else {
		if c.Mode != consts.BasicMode && c.ConfigPath == "" {
			logger.Errorf("config path is empty")
			param.Help()
			os.Exit(1)
		}
		if c.Mode == consts.RemoteMode && (c.RemoteAddr == "" || c.Secret == "") {
			logger.Errorf("remote addr or secret is empty")
			param.Help()
			os.Exit(1)
		}
	}
}

func (a *Agent) receiveClientId() {
	jwt, clientId, err := a.GetClientId()
	if err != nil {
		panic(err)
	}
	a.ClientId = clientId
	a.Jwt = jwt
	a.Header = map[string]string{
		"Authorization": jwt,
	}
	logger.Infof("Received client id: %s", clientId)
	logger.Infof("Received jwt: %s", jwt)
}
