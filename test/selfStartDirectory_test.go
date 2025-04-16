package test

import (
	"augeu-agent/internal/pkg/agent"
	"testing"
)

const (
	selfStartDirectoryRule = `
rule "selfStartDirectoryRule6" "basic"  salience 0
begin
	ret=fileSysUtils.LsFile("C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup")
	printer.PrintStrSlice(ret,"info")
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
