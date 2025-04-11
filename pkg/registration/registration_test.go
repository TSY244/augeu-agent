package registration

import (
	"golang.org/x/sys/windows/registry"
	"testing"
)

//func TestGetWindowsUserNames(t *testing.T) {
//	names, err := GetWindowsUserNamesByRegistry()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	for _, name := range names {
//		t.Log(name)
//	}
//}

func TestGetSubKeys(t *testing.T) {
	names, err := GetSubKeys(registry.LOCAL_MACHINE, IFEOPath)
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range names {
		t.Log(name)
	}
}

func TestGetDebugger(t *testing.T) {
	names, err := GetSubKeys(registry.LOCAL_MACHINE, IFEOPath)
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range names {
		debugger, err := GetDebuggerValue(name)
		if err != nil {
			//t.Error(err)
			continue
		}
		t.Log("name: ", name, "debugger: ", debugger)
	}
}
