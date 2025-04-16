package agent

import "testing"

const (
	StrUtilsCheck = `
rule "name testGetRegPathValueNames" "i can"  salience 0
begin
	value1="123"
	myArray = strUtils.SplitStr("123;345;456",";")
	println(myArray)
end
`
)

func TestReg(t *testing.T) {
	agentConf := Config{
		Mode: BasicMode,
	}
	agent := NewAgent(&agentConf)
	agent.Rule = `
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
	agent.Run()
}

func TestStrUtils(t *testing.T) {
	agentConf := Config{
		Mode: BasicMode,
	}
	agent := NewAgent(&agentConf)
	agent.Rule = StrUtilsCheck
	agent.Run()
}
