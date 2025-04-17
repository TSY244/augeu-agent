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

func TestGetPathSubKeys(t *testing.T) {
	names, err := GetPathSubKeys("HKEY_LOCAL_MACHINE\\HARDWARE")
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range names {
		t.Log(name)
	}
}

func TestGetRegPathValueNames(t *testing.T) {
	names, err := GetRegPathValueNames("HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Tpm")
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range names {
		t.Log(name)
	}
}

func TestGetRegPathValue(t *testing.T) {
	names, err := GetRegPathValueNames("HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows Mail")
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range names {
		value, err := GetRegPathValue("HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows Mail", name)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(name, value)
	}
}

func TestIsHavePath(t *testing.T) {
	isH := IsHavePath(`HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run`)
	if isH {
		t.Log("have path")
	}

}
