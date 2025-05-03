package windowsApi

import "testing"

func TestGetConnetions(t *testing.T) {
	tcpConnections, err := GetTCPConnections()
	if err != nil {
		t.Logf("GetTCPConnections err: %v", err)
		return
	}
	for _, conn := range tcpConnections {
		t.Logf("LocalIP: %v, LocalPort: %v, PID: %v, Protocol: %v, RemoteIP: %v, RemotePort: %v", conn.LocalIP, conn.LocalPort, conn.PID, conn.Protocol, conn.RemoteIP, conn.RemotePort)
	}
}
