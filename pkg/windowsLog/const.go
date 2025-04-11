package windowsLog

const (
	LoginEvenType       EventNameType = "LoginEvenType"
	RdpEventType        EventNameType = "RdpEventType"
	ServiceEventType    EventNameType = "ServiceType"
	CreateProcessType   EventNameType = "CreateProcessType"
	PowershellEventType EventNameType = "PowershellEventType"
	ReadLsassEventType  EventNameType = "ReadLsassEventType"
	SystemEventType     EventNameType = "SystemEventType"
	UserEventType       EventNameType = "UserEventType"
	SysmonEventType     EventNameType = "SysmonEventType"
	RegistryEventType   EventNameType = "RegistryEventType"
)

var (
	EventToFilePath = map[EventNameType][]string{
		LoginEvenType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		RdpEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		ServiceEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\System.evtx",
		},
		CreateProcessType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		PowershellEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Microsoft-Windows-PowerShell%4Operational.evtx",
		},
		ReadLsassEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		SystemEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\System.evtx",
		},
		UserEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		SysmonEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Microsoft-Windows-Sysmon%4Operational.evtx",
		},
		RegistryEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Microsoft-Windows-Sysmon%4Operational.evtx",
		},
	}
)

// runBase info key
const (
	MachineUUIDKey = "MachineUUID"
	EventIdKey     = "EventID"
	EventTimeKey   = "EventTime"
	DescriptionKey = "Description"
)

// login event info key
const (
	LoginTypeKey       = "LoginType"
	SourceIpKey        = "SourceIp"
	UsernameKey        = "Username"
	SubjectUsernameKey = "SubjectUsername"
	SubjectDomainKey   = "SubjectDomain"
	ProcessNameKey     = "ProcessName"
)

// login event value path
// 用于解析GoEvtxMap中的值，需要提供path 去解析map 中的value
const (
	usernamePath        = "/Event/EventData/TargetUserName"
	ipAddressPath       = "/Event/EventData/IpAddress"
	logonTypePath       = "/Event/EventData/LogonType"
	subjectUserNamePath = "/Event/EventData/SubjectUserName"
	subjectDomainPath   = "/Event/EventData/SubjectDomainName"
	processNamePath     = "/Event/EventData/ProcessName"
)

// rdp event info key
const (
	AccountNameKey   = "AccountName"
	AccountDomainKey = "AccountDomain"
	ClientNameKey    = "ClientName"
	ClientAddressKey = "ClientAddress"
)

// rdp event value path
const (
	AccountNamePath   = "/Event/EventData/AccountName"
	AccountDomainPath = "/Event/EventData/AccountDomain"
	ClientNamePath    = "/Event/EventData/ClientName"
	ClientAddressPath = "/Event/EventData/ClientAddress"
)

// service event info key
const (
	ServiceEventServiceNamePathKey = "ServiceName"
	ServiceEventImagePathKey       = "ImagePath"
	ServiceEventServiceTypePathKey = "ServiceType"
	ServiceEventStartTypePathKey   = "StartType"
	ServiceEventAccountNamePathKey = "AccountName"
)

// service event info key
const (
	ServiceEventServiceNamePath = "/Event/EventData/ServiceName"
	ServiceEventImagePath       = "/Event/EventData/ImagePath"
	ServiceEventServiceTypePath = "/Event/EventData/ServiceType"
	serviceEventStartTypePath   = "/Event/EventData/StartType"
	ServiceEventAccountNamePath = "/Event/EventData/AccountName"
)

// powershell event info key
const (
	PowershellEventScriptBlockTextKey = "ScriptBlockText"
	PowershellEventPath               = "Path"
)

// powershell event path key
const (
	PowershellEventScriptBlockTextPath = "/Event/EventData/ScriptBlockText"
	PowershellEventPathPath            = "/Event/EventData/Path"
)

// user event info key
const (
	UserEventTargetUserNameKey    = "TargetUserName"    // 新创建的用户的名称
	UserEventTargetDomainNameKey  = "TargetDomainName"  // 新创建的用户的域名
	UserEventTargetSidKey         = "TargetSid"         // 新创建的用户的SID
	UserEventSubjectUserNameKey   = "SubjectUserName"   // 执行操作的用户
	UserEventSubjectDomainNameKey = "SubjectDomainName" // 执行操作的用户的域名
	UserEventSubjectUserSidKey    = "SubjectUserSid"    // 执行操作的用户的SID
)

// user event path key
const (
	UserEventTargetUserNamePath    = "/Event/EventData/TargetUserName"
	UserEventTargetDomainNamePath  = "/Event/EventData/TargetDomainName"
	UserEventTargetSidPath         = "/Event/EventData/TargetSid"
	UserEventSubjectUserNamePath   = "/Event/EventData/SubjectUserName"
	UserEventSubjectDomainNamePath = "/Event/EventData/SubjectDomainName"
	UserEventSubjectUserSidPath    = "/Event/EventData/SubjectUserSid"
)
