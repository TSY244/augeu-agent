package dns

import (
	"fmt"
	"github.com/google/gopacket/pcap"
)

func GetAllDevices() ([]pcap.Interface, error) {
	return pcap.FindAllDevs()
}

func GetAllDevicesDetail() ([]string, error) {
	devices, err := GetAllDevices()
	if err != nil {
		return nil, err
	}
	ret := make([]string, len(devices))
	for i, v := range devices {
		ret[i] = fmt.Sprintf("接口名字：%s, 描述：%s，IP地址：%v", v.Name, v.Description, v.Addresses)
	}
	return ret, nil
}

func MonitorDNSInterface(interfaceName string) {

}
