package powershell

import "testing"

func TestRunExec(t *testing.T) {
	cmd := "powershell -c \"Get-Process\""
	result, err := RunExec(cmd)
	if err != nil {
		t.Errorf("RunExec failed: %v", err)
	}
	t.Logf("result: %v", result)

}

func TestGetScheduledTaskCommands(t *testing.T) {
	ret, err := GetScheduledTaskCommands()
	if err != nil {
		t.Errorf("GetScheduledTaskCommands failed: %v", err)
	} else {
		t.Logf("result: %v", ret)
	}
}

func TestGetBitsAdminInfo(t *testing.T) {
	ret, err := GetBitsAdminInfo()
	if err != nil {
		t.Errorf("GetBitsAdminInfo failed: %v", err)
	} else {
		t.Logf("result: %v", ret)
	}
}

func TestT(t *testing.T) {
	T()
}
