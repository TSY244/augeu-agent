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
	augeu := agent.NewAgent(&param.BaseConfig)
	augeu.Run()
}
