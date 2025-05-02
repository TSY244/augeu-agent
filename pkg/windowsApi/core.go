package windowsApi

import "syscall"

var (
	Kernel32                     = InitDll("kernel32.dll")
	procCreateToolhelp32Snapshot = GetFunction(Kernel32, procCreateToolhelp32SnapshotKey)
	procProcess32First           = GetFunction(Kernel32, procProcess32FirstKey)
	procProcess32Next            = GetFunction(Kernel32, procProcess32NextKey)
	procGetModuleFileNameEx      = GetFunction(Kernel32, procGetModuleFileNameExKey)
	procVirtualQueryEx           = GetFunction(Kernel32, procVirtualQueryExKey)
	procReadProcessMemory        = GetFunction(Kernel32, procReadProcessMemoryKey)
)

var (
	User32 = InitDll("user32.dll")
)

var (
	psapi                    = InitDll("psapi.dll")
	procGetModuleInformation = GetFunction(psapi, procGetModuleFileNameExKey)
)

func InitDll(dllName string) *syscall.LazyDLL {
	return syscall.NewLazyDLL(dllName)
}

func GetFunction(dll *syscall.LazyDLL, funcName string) *syscall.LazyProc {
	return dll.NewProc(funcName)
}
