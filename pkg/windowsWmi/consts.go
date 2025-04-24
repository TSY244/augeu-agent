package windowsWmi

const (
	QueryUuidKey          = "SELECT UUID FROM Win32_ComputerSystemProduct"
	QueryOsNameKey        = "SELECT Caption FROM Win32_OperatingSystem"
	QueryOsVersionKey     = "SELECT Version FROM Win32_OperatingSystem"
	QueryHotFixKey        = "SELECT Description, HotFixID, InstalledBy, InstalledOn FROM Win32_QuickFixEngineering"
	QueryUsersKey         = "SELECT Name, Description, LocalAccount, SID FROM Win32_UserAccount"
	QueryServiceKey       = "SELECT Name, DisplayName, State, StartMode, StartName FROM Win32_Service"
	QueryServiceDetailKey = "SELECT Name, DisplayName, State, StartMode, StartName, PathName FROM Win32_Service"
)
