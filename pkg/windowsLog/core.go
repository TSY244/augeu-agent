package windowsLog

import (
	"augeu/agent/pkg/windowsWmi"

	"augeu/public/pkg/logger"
	_const "augeu/public/util/const"
	"errors"
	"fmt"
	"github.com/0xrawsec/golang-evtx/evtx"
	"os"
)

type EventNameType string

type EventUnit map[string]interface{}
type ExternalFunctionForMapChan func(evtxMap chan *evtx.GoEvtxMap) error

type ExternalFunctionForMap func(evtxMap *evtx.GoEvtxMap)

var (
	FunctionMap = map[EventNameType]ExternalFunctionForMapChan{
		LoginEvenType:       loginEvent,
		RdpEventType:        rdpEvent,
		ServiceEventType:    serviceEvent,
		CreateProcessType:   createProcessEvent,
		PowershellEventType: powershellEvent,
		ReadLsassEventType:  readLsassEvent,
		SystemEventType:     systemEvent,
		UserEventType:       userEvent,
		RegistryEventType:   test,
	}
)

// -------------------------------------- public --------------------------------------

// Run 用于启动 eventName 对应的处理方式
//
// 参数：
//   - eventName: 事件名称
//   - mapChanFunctions: 外部函数列表，用于处理事件，默认使用 FunctionMap 中的函数，也可以通过 RegisterFunctionMap 注册新的函数
func Run(eventName EventNameType, mapChanFunctions ...ExternalFunctionForMapChan) error {
	eventFunc, ok := FunctionMap[eventName]
	if !ok {
		return errors.New("event not found")
	}
	if mapChanFunctions != nil {
		mapChanFunctions = append(mapChanFunctions, eventFunc)
	} else {
		mapChanFunctions = []ExternalFunctionForMapChan{
			eventFunc,
		}
	}
	return runBase(eventName, mapChanFunctions)
}

func RegisterFunctionMap(eventName EventNameType, function ExternalFunctionForMapChan) {
	FunctionMap[eventName] = function
}

func GetEventsForLoginEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	logger.Info("len(evtxMap): ", len(evtxMap))
	for event := range evtxMap {
		if _, ok := LoginEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		addLoginEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
		logger.Info("eventInfo: ", eventInfo, "len events: ", len(events), "event: ", event.EventID(), "eventInfo: ", eventInfo)
	}
	return events
}

func GetEventsForRdpEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := RdpEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		addRdpEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func GetEventsForServiceEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := ServiceEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		addServiceEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func GetEventsForCreateProcessEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := CreateProcessEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		//addCreateProcessEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func GetEventsForPowershellEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := PowershellEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		//addPowershellEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func GetEventsForReadLsassEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := ReadLsassEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		//addReadLsassEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func GetEventsForSystemEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := SystemEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		//addSystemEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func GetEventsForUserEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := UserEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		addUserEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

// base functions

func Wrapper(path string) *evtx.GoEvtxPath {
	p := evtx.Path(path)
	return &p
}

func GetString(evtxMap *evtx.GoEvtxMap, path *evtx.GoEvtxPath) string {
	value, err := evtxMap.GetString(path)
	if err != nil {
		logger.Error("GetString -> get string error: ", err)
		return ""
	}
	return value
}

// -------------------------------------- private --------------------------------------

func runBase(eventName EventNameType, mapChanFunctions []ExternalFunctionForMapChan) error {
	pathList, ok := EventToFilePath[eventName]
	if !ok {
		return errors.New("event file path not found")
	}
	for _, path := range pathList {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		eventReader, err := evtx.New(file)
		if err != nil {
			return err
		}
		eventMapChan := eventReader.FastEvents()
		for _, f := range mapChanFunctions {
			if err := f(eventMapChan); err != nil {
				return err
			}
		}

		// 资源释放
		file.Close()
		eventReader.Close()
	}

	return nil
}

func getBaseInfo(event *evtx.GoEvtxMap) EventUnit {
	// time : year-month-day hour:minute:second
	machineUUid, err := windowsWmi.QueryUuid()
	if err != nil {
		fmt.Println("getBaseInfo -> get windows guid error: ", err)
	}
	return EventUnit{
		EventIdKey:     event.EventID(),
		EventTimeKey:   event.TimeCreated().Format(_const.TimeFormat),
		MachineUUIDKey: machineUUid,
	}
}

// login event functions

// loginEvent 用于处理登录事件
func loginEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForLoginEvent(evtxMap)
	fmt.Println(events)
	return nil
}
func getLoginType(typeId string) string {
	switch typeId {
	case "2":
		return "Interactive（交互式登录）"
	case "3":
		return "Network（网络登录）"
	case "4":
		return "Batch（批处理登录）"
	case "5":
		return "Service（服务登录）"
	case "7":
		return "Unlock（解锁登录）"
	case "8":
		return "NetworkCleartext（网络明文登录）"
	case "9":
		return "NewCredentials（新凭证登录）"
	case "10":
		return "RemoteInteractive（远程交互登录）"
	case "11":
		return "CachedInteractive（缓存交互登录）"
	default:
		logger.Error("getLoginType -> login type not found: ", typeId)
		return "Unknown"
	}
}

func addLoginEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	(*eventUnit)[LoginTypeKey] = getLoginType(GetString(evtxMap, Wrapper(logonTypePath)))
	(*eventUnit)[UsernameKey] = GetString(evtxMap, Wrapper(usernamePath))
	(*eventUnit)[SourceIpKey] = GetString(evtxMap, Wrapper(ipAddressPath))
	(*eventUnit)[SubjectUsernameKey] = GetString(evtxMap, Wrapper(subjectUserNamePath))
	(*eventUnit)[SubjectDomainKey] = GetString(evtxMap, Wrapper(subjectDomainPath))
	(*eventUnit)[ProcessNameKey] = GetString(evtxMap, Wrapper(processNamePath))
}

// rdp event functions

// rdpEvent 用于处理远程桌面事件
func rdpEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForRdpEvent(evtxMap)
	fmt.Println(events)
	return nil
}

func addRdpEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	(*eventUnit)[AccountNameKey] = GetString(evtxMap, Wrapper(AccountNamePath))
	(*eventUnit)[AccountDomainKey] = GetString(evtxMap, Wrapper(AccountDomainPath))
	(*eventUnit)[ClientAddressKey] = GetString(evtxMap, Wrapper(ClientAddressPath))
	(*eventUnit)[ClientNameKey] = GetString(evtxMap, Wrapper(ClientNamePath))
}

// service event functions
// serviceEvent 用于处理服务事件
func serviceEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForServiceEvent(evtxMap)
	fmt.Println(events)
	return nil
}

func addServiceEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	(*eventUnit)[ServiceEventServiceNamePathKey] = GetString(evtxMap, Wrapper(ServiceEventServiceNamePath))
	(*eventUnit)[ServiceEventImagePathKey] = GetString(evtxMap, Wrapper(ServiceEventImagePath))
	(*eventUnit)[ServiceEventServiceTypePathKey] = GetString(evtxMap, Wrapper(ServiceEventServiceTypePath))
	(*eventUnit)[ServiceEventStartTypePathKey] = GetString(evtxMap, Wrapper(serviceEventStartTypePath))
	(*eventUnit)[ServiceEventAccountNamePathKey] = GetString(evtxMap, Wrapper(ServiceEventAccountNamePath))

}

// create process event functions
// createProcessEvent 用于处理进程创建事件
func createProcessEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForCreateProcessEvent(evtxMap)
	fmt.Println(events)
	return nil
}

// func addCreateProcessEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
// 	(*eventUnit)[CreateProcessEventProcessNamePathKey] = GetString(evtxMap, Wrapper(CreateProcessEventProcessNamePath))
//}

// powershell event functions
// powershellEvent 用于处理powershell事件
func powershellEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForPowershellEvent(evtxMap)
	fmt.Println(events)
	return nil
}

func addPowershellEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	//(*eventUnit)[PowershellEventCommandLinePathKey] = GetString(evtxMap, Wrapper(PowershellEventCommandLinePath))
}

// read lsass event functions
// readLsassEvent 用于处理读取lsass事件
func readLsassEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForReadLsassEvent(evtxMap)
	fmt.Println(events)
	return nil
}
func addReadLsassEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	//(*eventUnit)[ReadLsassEventProcessNamePathKey] = GetString(evtxMap, Wrapper(ReadLsassEventProcessNamePath))
}

// system event functions
// systemEvent 用于处理系统事件
func systemEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForSystemEvent(evtxMap)
	fmt.Println(events)
	return nil
}

func addSystemEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {

}

// user event functions
// userEvent 用于处理用户事件
func userEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := GetEventsForUserEvent(evtxMap)
	fmt.Println(events)
	return nil
}

func addUserEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	(*eventUnit)[UserEventTargetUserNameKey] = GetString(evtxMap, Wrapper(UserEventTargetUserNamePath))
	(*eventUnit)[UserEventTargetDomainNameKey] = GetString(evtxMap, Wrapper(UserEventTargetDomainNamePath))
	(*eventUnit)[UserEventTargetSidKey] = GetString(evtxMap, Wrapper(UserEventTargetSidPath))
	(*eventUnit)[UserEventSubjectUserNameKey] = GetString(evtxMap, Wrapper(UserEventSubjectUserNamePath))
	(*eventUnit)[UserEventSubjectDomainNameKey] = GetString(evtxMap, Wrapper(UserEventSubjectDomainNamePath))
	(*eventUnit)[UserEventSubjectUserSidKey] = GetString(evtxMap, Wrapper(UserEventSubjectUserSidPath))
	(*eventUnit)[DescriptionKey] = UserEvent[evtxMap.EventID()]
}

func test(evtxMap chan *evtx.GoEvtxMap) error {
	have := map[int64]interface{}{}
	for event := range evtxMap {
		//if _, ok := RegistryEvent[event.EventID()]; !ok {
		//	continue
		//}
		if _, ok := have[event.EventID()]; !ok {
			have[event.EventID()] = nil
			fmt.Println(event.EventID())
		}
	}
	return nil
}
