package windowsApi

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_VM_READ           = 0x0010
	MEM_COMMIT                = 0x00001000
	MEM_PRIVATE               = 0x00020000
	PAGE_READWRITE            = 0x04
)

type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

// SearchProcessMemory 在指定进程的内存中搜索关键字
func SearchProcessMemory(pid uint32, keyword string) ([]uintptr, error) {
	// 打开进程
	hProcess, err := openProcess(pid)
	if err != nil {
		return nil, fmt.Errorf("openProcess failed: %v", err)
	}
	defer syscall.CloseHandle(hProcess)

	// 准备搜索的关键字
	searchBytes := []byte(keyword)
	if len(searchBytes) == 0 {
		return nil, fmt.Errorf("empty keyword")
	}

	var results []uintptr
	var address uintptr

	// 遍历进程内存区域
	for {
		var mbi MEMORY_BASIC_INFORMATION
		ret, err := VirtualQueryEx(hProcess, address, unsafe.Pointer(&mbi), unsafe.Sizeof(mbi))
		if ret == 0 {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("VirtualQueryEx failed at 0x%x: %v", address, err)
		}

		// 只搜索已提交的、可读写的私有内存区域
		if mbi.State == MEM_COMMIT && mbi.Type == MEM_PRIVATE && mbi.Protect&PAGE_READWRITE != 0 {
			found, err := searchMemoryRegion(hProcess, mbi.BaseAddress, mbi.RegionSize, searchBytes)
			if err != nil {
				return nil, fmt.Errorf("searchMemoryRegion failed at 0x%x: %v", mbi.BaseAddress, err)
			}
			results = append(results, found...)
		}

		// 移动到下一个内存区域
		address = mbi.BaseAddress + mbi.RegionSize
	}

	return results, nil
}

// openProcess 打开指定PID的进程
func openProcess(pid uint32) (syscall.Handle, error) {
	handle, err := syscall.OpenProcess(PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, pid)
	if err != nil {
		return 0, err
	}
	return handle, nil
}

// searchMemoryRegion 在内存区域中搜索关键字
func searchMemoryRegion(hProcess syscall.Handle, baseAddress, regionSize uintptr, searchBytes []byte) ([]uintptr, error) {
	var found []uintptr

	// 读取内存块
	buffer := make([]byte, regionSize)
	var bytesRead uintptr
	success, err := ReadProcessMemory(hProcess, baseAddress, unsafe.Pointer(&buffer[0]), regionSize, &bytesRead)
	if !success || bytesRead == 0 {
		return nil, err
	}

	// 在读取的缓冲区中搜索关键字
	offset := 0
	for {
		idx := bytes.Index(buffer[offset:], searchBytes)
		if idx == -1 {
			break
		}
		foundAddr := baseAddress + uintptr(offset+idx)
		found = append(found, foundAddr)
		offset += idx + len(searchBytes)
	}

	return found, nil
}

// HexDump 生成内存区域的十六进制转储
func HexDump(pid uint32, address uintptr, size uint) (string, error) {
	hProcess, err := openProcess(pid)
	if err != nil {
		return "", fmt.Errorf("openProcess failed: %v", err)
	}
	defer syscall.CloseHandle(hProcess)

	buffer := make([]byte, size)
	var bytesRead uintptr
	success, err := ReadProcessMemory(hProcess, address, unsafe.Pointer(&buffer[0]), uintptr(size), &bytesRead)
	if !success || bytesRead == 0 {
		return "", err
	}

	return hex.Dump(buffer), nil
}
