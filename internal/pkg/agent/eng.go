package agent

import "augeu-agent/internal/pkg/engine/engUtils"

//func (a *Agent) AddObject() {
//	// 常规方法注册
//	a.Eng.InitObject("println", fmt.Println)
//
//	// 添加reg 注册对象
//	a.Eng.InitObject("reg", NewReg())
//}

func (a *Agent) SetEnv() {
	for k, v := range a.Conf.Env.GetEnvMap() {
		engUtils.CoreEnv.SetEnv(k, v)
	}
}
