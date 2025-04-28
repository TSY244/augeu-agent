# 使用



```
.\AugetAgent.exe -h
Usage of AugetAgent.exe:
  -cp string
        配置文件路径
  -mode string
        加载方式: basic, remote, local
                basic: 从程序内部加载默认规则
                remote: 加载远端提供的配置文件, -cp 参数需要提供一个远程url
                local: 从本地加载一个配置文件, -cp 需要提供一个本地文件路径
                 (default "basic")
  -r string
        server api 地址，请注意加上/api/v1，举个栗子：http://127.0.0.1/api/v1
  -s string
        server 提供的 secret
  -wbapi string
        微步api密钥
  -ws string
        server 提供的websocket 地址，注意使用ws://开头
```

## 本地模式

本地模式启动

本地模式使用的是程序内置的规则

直接运行即可

```
.\AugetAgent.exe -wbapi  XXXXXXXXXXXXXX
```

或者是使用`-mode basic` 参数进行指定

![image-20250428221210288](img\image-20250428221059918.png)



## 远程文件模式

使用python 起一个服务，把文件挂载，然后使用`-mode remoteFile` 的模式进行加载

```
.\AugetAgent.exe -mode remoteFile -cp http://127.0.0.1:8000/rule.txt -wbapi  XXXXXXXXXXXXXX
```



## 远程api 模式

先启动backend

![image-20250428223854304](D:\code\GoLand\augeu-agent\img\image-20250428223854304.png)

这个token 就是secrete

```
.\AugetAgent.exe -mode remoteApi -r http://127.0.0.1:8080/api/v1 -s R1gmbk9yPiNEY3RYNVdSd -wbapi  XXXXXXXXXXXXXX
```



# 日志输出分析

## rule 内部日志

### 标记

每一个rule 内部输出的日志

> [task log] [REMIND] [自启动文件夹下检测1] desktop.ini

带有前缀 `[task log] `

![image-20250428224944255](D:\code\GoLand\augeu-agent\img\image-20250428224944255.png)

## 程序日志

### 格式

采用格式

```
时间戳 level error
2025-04-28T22:47:48+08:00 ERR get hash error: open explorer.exe: The system cannot find the file specified.
```

![image-20250428224909131](D:\code\GoLand\augeu-agent\img\image-20250428224909131.png)

### 常见日志解析

1. reg err

   ``` 
   2025-04-28T22:47:48+08:00 ERR get hash error: open explorer.exe: The system cannot find the file specified.
   ```

   这个日志输出表示没有这个注册表路径，有可能是错误。有可能是真没有危害



# 扩展

## 代码上扩展

### 添加模块

rule 中的代码主要是使用模板和反射进行加载

所以只支持写好的规则，然后进行注册，最后在rule 中使用

internal/pkg/agent/core.go

```go

func NewAgent(c *param.Config) *Agent {
	checkConf(c)
	rootCtx, cancel := context.WithCancel(context.Background())
	
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
		"check":        engUtils.NewCheck(),
		"service":      engUtils.NewService(),
		"err":          consts.ErrKey,
		"schedule":     engUtils.NewSchedule(),
		"phs":          engUtils.NewPhs(),
	}
	return agent
}
```

在`agent.ApiOuter` 这个map 中进行注册

key 为rule 中调用的时候的名字

value 是函数指针，或者是结构体指针。建议放结构体，然后函数是成员函数的形式。

举个例子，比如reg

internal/pkg/engine/engUtils/reg.go

```go
package engUtils

import (
	"augeu-agent/internal/utils/consts"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/registration"
	"golang.org/x/sys/windows/registry"
	"regexp"
)

type Reg struct {
}

func NewReg() *Reg {
	return &Reg{}
}

func (r *Reg) GetPathSubKeys(path string) []string {
	names, err := registration.GetPathSubKeys(path)
	if err != nil {
		logger.Errorf("get path sub keys error: %v", err)
		return nil
	}
	return names
}

func (r *Reg) GetRegPathValueNames(path string) []string {
	//return registration.GetRegPathValueNames(path)
	names, err := registration.GetRegPathValueNames(path)
	if err != nil {
		logger.Errorf("get reg path value names error: %v", err)
		return nil
	}
	return names
}

func (r *Reg) GetRegPathValue(path string, name string) string {
	//return registration.GetRegPathStringValue(path, name)
	value, err := registration.GetRegPathStringValue(path, name)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return value
}

func (r *Reg) RegNameType(reg registry.Key, name string) uint32 {
	//return registration.RegNameType(reg, name)
	ret, err := registration.RegNameType(reg, name)
	if err != nil {
		logger.Errorf("get reg name type error: %v", err)
		return 0
	}
	return ret
}

func (r *Reg) IsPathWithSubKey(path string, subKey string) bool {
	return registration.IsPathWithSubKey(path, subKey)
}
func (r *Reg) IsPathWithName(path string, name string) bool {
	return registration.IsPathWithName(path, name)
}

func (r *Reg) GetPathFromCmd(cmd string) string {
	re := regexp.MustCompile(`^"([^"]+)"|^([^ "]+)`) // 匹配带引号和不带引号的路径
	matches := re.FindStringSubmatch(cmd)
	if len(matches) == 0 {
		return ""
	}

	var path string
	if matches[1] != "" { // 带引号的路径
		path = matches[1]
	} else { // 不带引号的路径
		path = matches[2]
	}
	return path
}

func (r *Reg) IsHavePath(path string) bool {
	return registration.IsHavePath(path)
}

func (r *Reg) GetDefaultRegPathValue(path string) string {
	ret, err := registration.GetDefaultRegPathValue(path)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return ret
}

```



### 模块介绍



