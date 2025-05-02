package windowsApi

// function names

const (
	procCreateToolhelp32SnapshotKey = "CreateToolhelp32Snapshot"
	procProcess32FirstKey           = "Process32FirstW"
	procProcess32NextKey            = "Process32NextW"
	procGetModuleFileNameExKey      = "GetModuleFileNameExW"
	procVirtualQueryExKey           = "VirtualQueryEx"
	procReadProcessMemoryKey        = "ReadProcessMemory"
)

const (
	TH32CS_SNAPPROCESS = 0x00000002
)

// errors
const (
	INVALID_HANDLE_VALUE = 0
)
