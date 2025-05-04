package netstat

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/windowsApi"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type ProcessFunc func([]net.IP)
type FilterFunc func(conn windowsApi.ConnectionInfo) bool

type ConnCache struct {
	entries map[uint64]*net.IP
	sync.RWMutex
}

var (
	Cache      *ConnCache
	CacheFlag  bool
	CurrentPid int
)

func init() {
	if Cache == nil {
		Cache = &ConnCache{
			entries: make(map[uint64]*net.IP),
		}
	}
	CacheFlag = true
	CurrentPid = os.Getpid()
}

func Monitor(target string, fs ...ProcessFunc) error {
	ips := NewIpTable()
	if target == "" {
		return fmt.Errorf("target is empty")
	}
	if net.ParseIP(target) == nil {
		err := ips.UpdateIPs(target)
		if err != nil {
			return err
		}
		go ips.StartDNSUpdater(target)
	} else {
		ips.Lock()
		ips.IPs = []net.IP{net.ParseIP(target)}
		ips.Unlock()
	}
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()
	for range ticker.C {
		checkConnection(ips, fs...)
	}
	return nil
}

func MonitorTcpAndUdp(target string, isStrong bool) error {
	return Monitor(target, ProcessTcpFunc(isStrong), ProcessUdpFunc(isStrong))
}

func MonitorDns(target string, isStrong bool) error {
	return Monitor(target, ProcessDnsFunc(isStrong))
}

func checkConnection(targetIPs *IpTable, fs ...ProcessFunc) {
	targetIPs.RLock()
	ips := make([]net.IP, len(targetIPs.IPs))
	copy(ips, targetIPs.IPs)
	targetIPs.RUnlock()

	var wg sync.WaitGroup
	wg.Add(len(fs))
	for _, f := range fs {
		go func() {
			defer wg.Done()
			f(ips)
		}()
	}
	//var tcpConnection, udpConnection []windowsApi.ConnectionInfo
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	tcpConnections, _ := windowsApi.GetTCPConnections()
	//	tcpConnection = tcpConnections
	//}()
	//go func() {
	//	defer wg.Done()
	//	udpConnections, _ := windowsApi.GetUDPConnections()
	//	udpConnection = udpConnections
	//}()
	//wg.Wait()
	//ProcessConnection(tcpConnection, ips)
	//ProcessConnection(udpConnection, ips)
}

func ProcessTcpFunc(isStrong bool) ProcessFunc {
	return func(ips []net.IP) {
		tcpConnections, _ := windowsApi.GetTCPConnections()
		ProcessConnection(tcpConnections, isStrong, ips)
	}
}

func ProcessUdpFunc(isStrong bool) ProcessFunc {
	return func(ips []net.IP) {
		udpConnections, _ := windowsApi.GetUDPConnections()
		ProcessConnection(udpConnections, isStrong, ips)
	}
}

func ProcessDnsFunc(isStrong bool) ProcessFunc {
	return func(ips []net.IP) {
		udpConnections, _ := windowsApi.GetUDPConnections()
		ProcessConnectionTestFunc(udpConnections, isStrong, ips)
		//ProcessConnection(udpConnections, isStrong, ips, func(conn windowsApi.ConnectionInfo) bool {
		//	if conn.LocalPort == 53 {
		//		return true
		//	}
		//	return false
		//})
	}
}

func ProcessConnection(conns []windowsApi.ConnectionInfo, isStrong bool, targets []net.IP, filter ...FilterFunc) {
	// 将目标IP转换为集合，实现O(1)查找
	targetSet := make(map[string]*net.IP, len(targets))
	for _, ip := range targets {
		targetSet[ip.String()] = &ip
	}

	for _, conn := range conns {
		for _, f := range filter {
			if !f(conn) {
				continue
			}
		}

		if conn.PID == uint32(CurrentPid) {
			continue
		}
		//if conn.RemoteIP == nil {
		//	continue
		//}
		remoteIP := conn.RemoteIP.String()
		if _, exists := targetSet[remoteIP]; !exists {
			continue
		}

		fingerprint := windowsApi.ConnectionFingerprint(conn)

		Cache.Lock()
		if !isStrong {
			if _, exists := Cache.entries[fingerprint]; exists {
				Cache.Unlock()
				continue
			}
		}

		logger.Infof("[检测到连接] PID:%d 远程地址:%s:%d 类型:%s",
			conn.PID,
			conn.RemoteIP,
			conn.RemotePort,
			conn.Protocol,
		)

		if CacheFlag {
			// 存储时间戳或空结构即可，这里保持原逻辑但修复指针问题
			Cache.entries[fingerprint] = targetSet[remoteIP]
		}
		Cache.Unlock()
	}
}
func SetCacheFlag(flag bool) {
	CacheFlag = flag
}

func ProcessConnectionTestFunc(conns []windowsApi.ConnectionInfo, isStrong bool, targets []net.IP, filter ...FilterFunc) {
	// 将目标IP转换为集合，实现O(1)查找
	targetSet := make(map[string]*net.IP, len(targets))
	for _, ip := range targets {
		targetSet[ip.String()] = &ip
	}

	for _, conn := range conns {
		for _, f := range filter {
			if !f(conn) {
				continue
			}
		}

		if conn.PID == uint32(CurrentPid) {
			continue
		}
		//if conn.RemoteIP == nil {
		//	continue
		//}
		remoteIP := conn.RemoteIP.String()
		//if _, exists := targetSet[remoteIP]; !exists {
		//	continue
		//}

		fingerprint := windowsApi.ConnectionFingerprint(conn)

		Cache.Lock()
		if !isStrong {
			if _, exists := Cache.entries[fingerprint]; exists {
				Cache.Unlock()
				continue
			}
		}

		logger.Infof("[检测到连接] PID:%d 远程地址:%s:%d 类型:%s",
			conn.PID,
			conn.RemoteIP,
			conn.RemotePort,
			conn.Protocol,
		)

		if CacheFlag {
			// 存储时间戳或空结构即可，这里保持原逻辑但修复指针问题
			Cache.entries[fingerprint] = targetSet[remoteIP]
		}
		Cache.Unlock()
	}
}
