package test

import (
	"augeu-agent/internal/pkg/agent"
	"augeu-agent/internal/pkg/param"
	"augeu-agent/internal/utils/consts"
	"testing"
)

const (
	RegSelfStartRule1 = `
rule "注册表自启动检测1" "basic"  salience 0
begin
	path="HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run"
	names=reg.GetRegPathValueNames(path)
	size=base.SizeForStr(names)
	forRange i:=names{
		ret=reg.GetRegPathValue(path,names[i])
		printer.Info(ret)		
		r=reg.GetPathFromCmd(ret)
		hash=fileSysUtils.GetHashWithFilePath(r)
		printer.Info(hash)
	}
end
`
)

func TestRegSelfStart(t *testing.T) {
	agentConf := param.Config{
		Mode: consts.BasicMode,
	}
	a := agent.NewAgent(&agentConf)
	a.Rule = RegSelfStartRule1
	a.Run()
}
