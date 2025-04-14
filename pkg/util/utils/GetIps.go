package utils

import "net"

func GetIps() (*[]string, error) {
	var ips []string

	adders, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range adders {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return &ips, nil
}
