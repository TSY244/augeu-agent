package engUtils

import "testing"

func TestPrinter(t *testing.T) {
	r := NewPrinter()
	r.Info("info")
	r.Warn("warn")
	r.Debug("debug")
}
