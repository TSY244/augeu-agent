package main

import (
	"augeu-agent/internal/pkg/agent"
	"augeu-agent/internal/pkg/param"
)

const (
	rule = `
rule "name testGetPathSubKeys" "i can"  salience 0
begin
	names = reg.GetPathSubKeys("HKEY_LOCAL_MACHINE\HARDWARE")
	println(names)
	println(@name)
end
rule "name testGetRegPathValueNames" "i can"  salience 0
begin
	names = reg.GetRegPathValueNames("HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Tpm")
	println(names)
	println("asdf"+ @name) 
end
`
)

func main() {
	param.Init()
	c := agent.Config{
		ConfigPath: param.BaseConfig.ConfigPath,
		Mode:       param.BaseConfig.Mode,
	}
	augeu := agent.NewAgent(&c)

	augeu.SetRule(rule)
	augeu.Run()
}
