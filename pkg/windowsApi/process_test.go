package windowsApi

import "testing"

//func TestGetAllProcess(t *testing.T) {
//	processes, err := GetAllProcess()
//	if err != nil {
//		t.Errorf("GetAllProcess failed: %v", err)
//	}
//	for _, process := range processes {
//		t.Logf("Process Name: %s, Pid: %d", process.Name, process.Pid)
//	}
//}

func TestGetExePath(t *testing.T) {
	processes, err := GetAllProcess()
	if err != nil {
		t.Errorf("GetAllProcess failed: %v", err)
	}
	for _, process := range processes {
		exePath, err := GetExePath(process.Pid)
		if err != nil {
			t.Errorf("GetExePath failed: %v", err)
		}
		t.Logf("Process Name: %s, Pid: %d, ExePath: %s", process.Name, process.Pid, exePath)
	}
}

func TestSearch(t *testing.T) {
	processes, err := GetAllProcess()
	if err != nil {
		t.Errorf("GetAllProcess failed: %v", err)
	}
	for _, process := range processes {
		hProcess, err := OpenProcess(process.Pid)
		if err != nil {
			t.Errorf("OpenProcess failed: %v", err)
		}

		ret, err := ScanProcessMemory(hProcess, "123&testaugust..redxcftgvyhbjnkmaskldjfhalks;jdf290318u7498uiojhsadnfjhbnaslkdjfh123&testaugust..redxcftgvyhbjnkmaskldjfhalks;jdf290318u7498uiojhsadnfjhbnaslkdjfh123&testaugust..redxcftgvyhbjnkmaskldjfhalks;jdf290318u7498uiojhsadnfjhbnaslkdjfh")
		if err != nil {
			t.Errorf("SearchProcessMem failed: %v", err)
		}
		if ret {
			exePath, err := GetExePath(process.Pid)
			if err != nil {
				t.Errorf("GetExePath failed: %v", err)
			}
			t.Logf("Process Name: %s, Pid: %d, ExePath: %s", process.Name, process.Pid, exePath)
			//t.Logf("Process Name: %s, Pid: %d, Found", process.Name, process.Pid)
		}
	}
}
