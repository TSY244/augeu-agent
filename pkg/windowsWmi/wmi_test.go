package windowsWmi

import "testing"

//const (
//	QueryUuidKey      = "SELECT UUID FROM Win32_ComputerSystemProduct"
//	QueryOsNameKey    = "SELECT Caption FROM Win32_OperatingSystem"
//	QueryOsVersionKey = "SELECT Version FROM Win32_OperatingSystem"
//	QueryHotFixKey    = "SELECT Description, HotFixID, InstalledBy, InstalledOn FROM Win32_QuickFixEngineering"
//)

func TestQuery(t *testing.T) {
	//ret, err := QueryUuid()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log("uuid: ", ret)
	//ret, err = QueryOsName()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log("osName: ", ret)
	//ret, err = QueryOsVersion()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log("osVersion: ", ret)
	//ret2, err := QueryHotFix()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log("hotFix: ", ret2)
	//ret3, err := QueryUsers()
	//if err != nil {
	//	t.Error(err)
	//}
	//for _, user := range ret3 {
	//	t.Log(user)
	//}
	//ret4, err := QueryScheduledTasks()
	//if err != nil {
	//	t.Error(err)
	//}
	//for _, task := range ret4 {
	//	t.Log(task)
	//}
	//ret5, err := QueryServices()
	//if err != nil {
	//	t.Error(err)
	//}
	//for _, task := range ret5 {
	//	t.Log(task)
	//}

	//ret6, err := QueryServicesDetail()
	//if err != nil {
	//	t.Error(err)
	//}
	//for _, task := range ret6 {
	//	t.Logf("name: %s, displayName: %s, state: %s, startMode: %s, startName: %s, pathName: %s",
	//		task.Name, task.DisplayName, task.State, task.StartMode, task.StartName, task.PathName)
	//}

	ret7, err := QueryScheduledTasks()
	if err != nil {
		t.Error(err)
		return
	}
	for _, r := range ret7 {
		//t.Logf("name: %s, author: %s, state: %s, description: %s, path: %s, uri: %s",
		//	r.TaskName, r.Author, r.State, r.Description, r.TaskPath, r.URI)
		t.Logf("path: %s", r.TaskPath)
	}
}

//
//type win32_ScheduledTask struct {
//	TaskName    string `wmi:"TaskName"`
//	Author      string `wmi:"Author"`
//	State       int    `wmi:"State"`
//	Description string `wmi:"Description"`
//	TaskPath    string `wmi:"TaskPath"`
//	URI         string `wmi:"URI"`
//}
