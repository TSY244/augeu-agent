package windowsLog

import "testing"

func TestRun(t *testing.T) {
	if err := Run(RegistryEventType); err != nil {
		t.Error(err)
	}
}
