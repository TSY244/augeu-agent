package windowsApi

// function names

const (
	procCreateToolhelp32SnapshotKey = "CreateToolhelp32Snapshot"
	procProcess32FirstKey           = "Process32FirstW"
	procProcess32NextKey            = "Process32NextW"
	procGetModuleFileNameExKey      = "GetModuleFileNameExW"
	procVirtualQueryExKey           = "VirtualQueryEx"
	procReadProcessMemoryKey        = "ReadProcessMemory"
	procGetExtendedTcpTableKey      = "GetExtendedTcpTable"
	procGetExtendedUdpTableKey      = "GetExtendedUdpTable"
)

const (
	TH32CS_SNAPPROCESS = 0x00000002
)

// errors
const (
	INVALID_HANDLE_VALUE = 0
)

const (
	TCP_TABLE_OWNER_PID_ALL = 5
	UDP_TABLE_OWNER_PID     = 1
)
