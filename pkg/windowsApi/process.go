package windowsApi

import (
	"bytes"
	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func GetAllProcess() ([]Process, error) {
	var processList []Process
	hSnapshot, _, err := procCreateToolhelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	if hSnapshot == INVALID_HANDLE_VALUE {
		return nil, err
	}
	defer syscall.CloseHandle(syscall.Handle(hSnapshot))
	var pe32 PROCESSENTRY32
	pe32.dwSize = uint32(unsafe.Sizeof(pe32))
	ret, _, err := procProcess32First.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return nil, err
	}
	for {
		processList = append(processList, Process{
			Name: syscall.UTF16ToString(pe32.szExeFile[:]),
			Pid:  pe32.th32ProcessID,
		})
		ret, _, err := procProcess32Next.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
		if err != nil && err.Error() != "The operation completed successfully." {
			return nil, err
		}
	}
	return processList, nil
}

func GetExePath(pid uint32) (string, error) {
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return "", err
	}

	return p.Exe()
}

func VirtualQueryEx(hProcess syscall.Handle, lpAddress uintptr, ipBuff unsafe.Pointer, deLength uintptr) (int, error) {
	ret, _, err := procVirtualQueryEx.Call(
		uintptr(hProcess),
		lpAddress,
		uintptr(ipBuff),
		deLength,
	)
	if ret == 0 {
		return 0, err
	}
	return int(ret), nil
}

func ReadProcessMemory(hProcess syscall.Handle, lpBaseAddress uintptr, lpBuffer unsafe.Pointer, nSize uintptr, lpNumberOfBytesRead *uintptr) (bool, error) {
	ret, _, err := procReadProcessMemory.Call(uintptr(hProcess), lpBaseAddress, uintptr(lpBuffer), nSize, uintptr(unsafe.Pointer(lpNumberOfBytesRead)))
	if ret == 0 {
		return false, err
	}
	return true, nil
}

func ScanProcessMemory(hProcess windows.Handle, targetStr string) (bool, error) {
	targetBytes := []byte(targetStr)
	if len(targetBytes) == 0 {
		return false, nil
	}

	var addr uintptr
	targetLen := len(targetBytes)
	var mbi windows.MemoryBasicInformation

	for {
		err := windows.VirtualQueryEx(hProcess, addr, &mbi, uintptr(unsafe.Sizeof(mbi)))
		if err != nil || mbi.RegionSize == 0 {
			break
		}

		if isReadable(mbi.Protect) {
			baseAddr := mbi.BaseAddress
			regionSize := mbi.RegionSize
			var prevRemaining []byte

			for offset := uintptr(0); offset < regionSize; {
				chunkSize := uintptr(1024 * 1024) // 1MB chunks
				if remaining := regionSize - offset; remaining < chunkSize {
					chunkSize = remaining
				}

				buffer := make([]byte, chunkSize)
				var bytesRead uintptr
				err := windows.ReadProcessMemory(hProcess, baseAddr+offset, &buffer[0], chunkSize, &bytesRead)
				if err != nil || bytesRead == 0 {
					break
				}

				data := buffer[:bytesRead]
				if len(prevRemaining) > 0 {
					data = append(prevRemaining, data...)
					prevRemaining = nil
				}

				if bytes.Contains(data, targetBytes) {
					return true, nil
				}

				// 保存可能跨块的尾部数据
				if len(data) >= targetLen {
					prevRemaining = data[len(data)-targetLen+1:]
				} else {
					prevRemaining = data
				}

				offset += chunkSize
			}
		}

		addr = mbi.BaseAddress + mbi.RegionSize
	}

	return false, nil
}

func isReadable(protect uint32) bool {
	return protect&(windows.PAGE_READONLY|
		windows.PAGE_READWRITE|
		windows.PAGE_EXECUTE_READ|
		windows.PAGE_EXECUTE_READWRITE|
		windows.PAGE_WRITECOPY|
		windows.PAGE_EXECUTE_WRITECOPY) != 0
}

func OpenProcess(pid uint32) (windows.Handle, error) {
	hProcess, err := windows.OpenProcess(PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, pid)
	if err != nil {
		return 0, err
	}
	return hProcess, nil
}

func CloseHandle(hProcess windows.Handle) error {
	err := windows.CloseHandle(hProcess)
	if err != nil {
		return err
	}
	return nil
}
