package registration

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
