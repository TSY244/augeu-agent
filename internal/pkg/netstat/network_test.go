package netstat

import "testing"

func TestMonitor(t *testing.T) {
	MonitorTcpAndUdp("win.au9u5t.fun", false)
}

func TestMonitorDns(t *testing.T) {
	MonitorDns("win.au9u5t.fun", false)
}
