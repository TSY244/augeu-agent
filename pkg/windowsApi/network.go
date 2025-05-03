package windowsApi

import (
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

type ConnectionInfo struct {
	LocalIP    net.IP
	LocalPort  uint16
	RemoteIP   net.IP
	RemotePort uint16
	PID        uint32
	Protocol   string
}

func GetTCPConnections() ([]ConnectionInfo, error) {
	return getConnections(procGetExtendedTcpTable, TCP_TABLE_OWNER_PID_ALL, "tcp")
}

func GetUDPConnections() ([]ConnectionInfo, error) {
	return getConnections(procGetExtendedUdpTable, UDP_TABLE_OWNER_PID, "udp")
}

func getConnections(proc *syscall.LazyProc, tableClass int, protocol string) ([]ConnectionInfo, error) {
	var size uint32
	rc, _, _ := proc.Call(
		0,
		uintptr(unsafe.Pointer(&size)),
		0,
		syscall.AF_INET,
		uintptr(tableClass),
		0,
	)
	//if err != nil {
	//	return nil, err
	//}
	if rc != 0 && rc != 122 {
		return nil, fmt.Errorf("get table size is err: %X", rc)
	}
	buffer := make([]byte, size)
	rc, _, _ = proc.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(&size)),
		0,
		syscall.AF_INET,
		uintptr(tableClass),
		0,
	)
	if rc != 0 {
		return nil, fmt.Errorf("get table is err: %X", rc)
	}
	numRows := *(*uint32)(unsafe.Pointer(&buffer[0]))
	if numRows == 0 {
		return nil, nil
	}
	var rowSize uint32
	switch tableClass {
	case TCP_TABLE_OWNER_PID_ALL:
		rowSize = 24
	case UDP_TABLE_OWNER_PID:
		rowSize = 12
	default:
		return nil, fmt.Errorf("unknown table class: %d", tableClass)
	}
	var connections []ConnectionInfo
	var offset uint32
	offset = 4
	for i := 0; i < int(numRows); i++ {
		conn := parseRow(buffer[offset:offset+rowSize], tableClass, protocol)
		connections = append(connections, conn)
		offset += rowSize
	}
	return connections, nil
}

func parseRow(row []byte, tableClass int, protocol string) ConnectionInfo {
	parseIP := func(addr uint32) net.IP {
		ip := make(net.IP, 4)
		binary.LittleEndian.PutUint32(ip, addr)
		return ip
	}

	parsePort := func(portBytes []byte) uint16 {
		if len(portBytes) < 2 {
			return 0
		}
		return binary.BigEndian.Uint16(portBytes[:2])
	}

	switch tableClass {
	case TCP_TABLE_OWNER_PID_ALL:
		return ConnectionInfo{
			LocalIP:    parseIP(binary.LittleEndian.Uint32(row[4:8])),
			LocalPort:  parsePort(row[8:10]),
			RemoteIP:   parseIP(binary.LittleEndian.Uint32(row[12:16])),
			RemotePort: parsePort(row[16:18]),
			PID:        binary.LittleEndian.Uint32(row[20:24]),
			Protocol:   protocol,
		}
	case UDP_TABLE_OWNER_PID:
		return ConnectionInfo{
			LocalIP:   parseIP(binary.LittleEndian.Uint32(row[0:4])),
			LocalPort: parsePort(row[4:6]),
			PID:       binary.LittleEndian.Uint32(row[8:12]),
			Protocol:  protocol,
		}
	default:
		return ConnectionInfo{Protocol: protocol}
	}
}

func ConnectionFingerprint(c ConnectionInfo) uint64 {
	return uint64(c.PID)<<32 | uint64(c.RemotePort)<<16 | uint64(c.LocalPort)
}
