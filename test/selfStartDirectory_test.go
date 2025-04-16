package test

import (
	"augeu-agent/internal/pkg/agent"
	"testing"
)

const (
	selfStartDirectoryRule = `
rule "selfStartDirectoryRule" "basic"  salience 0
begin
	ret=fileSysUtils.LsFile("C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup")
	printer.Warn(ret)
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
