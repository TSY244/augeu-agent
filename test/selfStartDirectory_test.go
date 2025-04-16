package test

import (
	"augeu-agent/internal/pkg/agent"
	"testing"
)

const (
	selfStartDirectoryRule = `
rule "自启动文件夹下检测1" "basic"  salience 0
begin
	path="C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup"
	files=fileSysUtils.LsFile(path)
	printer.Warn(@name+" 当前windows 不支持link 解析，请手动分析，当前路径 "+path+" 路径下有一下的文件：")
	printer.PrintStrSlice(files,"remind" ,@name)
	return false
end
rule "自启动文件夹下检测2" "basic"  salience 0
begin
	path="%appdata%\Microsoft\Windows\Start Menu\Programs\Startup"
	files=fileSysUtils.LsFile(path)
	printer.Warn(@name+" 当前windows 不支持link 解析，请手动分析，当前路径 "+path+" 路径下有一下的文件：")
	printer.PrintStrSlice(files,"remind",@name)
	return false
end
`
)

func TestRun(t *testing.T) {
	agentConf := agent.Config{
		Mode: agent.BasicMode,
	}
	a := agent.NewAgent(&agentConf)
	a.Rule = selfStartDirectoryRule
	a.Run()
}
