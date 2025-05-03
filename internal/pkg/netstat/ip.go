package netstat

import (
	"augeu-agent/pkg/logger"
	"net"
	"sync"
	"time"
)

type IpTable struct {
	sync.RWMutex
	IPs []net.IP
}

func NewIpTable() *IpTable {
	return &IpTable{
		IPs: []net.IP{},
	}
}

func (i *IpTable) UpdateIPs(target string) error {
	ips, err := net.LookupIP(target)
	if err != nil {
		logger.Errorf("DNS查询失败: %v", err)
		return err
	}

	i.Lock()
	defer i.Unlock()
	i.IPs = ips
	return err
}

func (i *IpTable) StartDNSUpdater(target string) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		err := i.UpdateIPs(target)
		if err != nil {
			logger.Errorf("DNS更新失败: %v", err)
		}
	}
}
