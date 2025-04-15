package registration

import "golang.org/x/sys/windows/registry"

// local registration
const (
	GuidKeyPath = "SOFTWARE\\Microsoft\\Cryptography"
	GuidKeyName = "MachineGuid"

	// UserNamesPath 对应的key
	UserNamesPath = "SAM\\SAM\\Domains\\Account\\Users\\Namess"
	// IFEO
	IFEOPath            = "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Image File Execution Options"
	IFEODebuggerKeyName = "Debugger"
)

var (
	RootPathMap = map[string]registry.Key{
		"HKLM":                registry.LOCAL_MACHINE,
		"HKEY_LOCAL_MACHINE":  registry.LOCAL_MACHINE,
		"HKCU":                registry.CURRENT_USER,
		"HKEY_CURRENT_USER":   registry.CURRENT_USER,
		"HKCR":                registry.CLASSES_ROOT,
		"HKEY_CLASSES_ROOT":   registry.CLASSES_ROOT,
		"HKEY_CURRENT_CONFIG": registry.CURRENT_CONFIG,
		"HKCC":                registry.CURRENT_CONFIG,
	}
)

//
//二进制(REG_BINARY)
//在注册表中，二进制是没有长度限制的，可以是任意个字节的长度。
//
//DWORD值(REG_DWORD)
//DWORD值是一个32位（4个字节，即双字）长度的整数。在注册表编辑器中，系统以十六进制的方式显示DWORD值。
//
//字符串值(REG_SZ)
//在注册表中，字符串值一般用来表示文件的描述、硬件的标识等，通常它是以空字符(\0)结尾的字符串。
//
//QWORD值(REG_QWORD)
//DWORD值是一个64位（8个字节，即四字）长度的数值。在注册表编辑器中，系统以十六进制的方式显示QWORD值。
//
//多字符串值(REG_MULTI_SZ)
//由两个空字符终止的空终止字符串数组。
//
//可扩充字符串值(REG_EXPAND_SZ)
//包含对环境变量的未扩展引用的空终止字符串（例如，“%PATH%”）。

const (
	REG_BINARY    = registry.BINARY
	REG_DWORD     = registry.DWORD
	REG_SZ        = registry.SZ
	REG_QWORD     = registry.QWORD
	REG_MULTI_SZ  = registry.MULTI_SZ
	REG_EXPAND_SZ = registry.EXPAND_SZ
)
