package convert

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/swaggerCore/models"
	"augeu-agent/pkg/util/convert"
	"augeu-agent/pkg/windowsLog"
	"augeu-agent/pkg/windowsWmi"
)

func LoginEvent2RLoginEventResq(event windowsLog.EventUnit) *models.LoginEvent {

	eventId := event[windowsLog.EventIdKey].(int64)
	eventTime := event[windowsLog.EventTimeKey].(string)
	machineUuid := event[windowsLog.MachineUUIDKey].(string)
	loginType := event[windowsLog.LoginTypeKey].(string)
	username := event[windowsLog.UsernameKey].(string)
	ipAddress := event[windowsLog.SourceIpKey].(string)
	subjectUserName := event[windowsLog.SubjectUsernameKey].(string)
	subjectDomain := event[windowsLog.SubjectDomainKey].(string)
	processName := event[windowsLog.ProcessNameKey].(string)

	// str -> *strfmt.DateTime
	tempTime, err := convert.StrTime2DateTime(eventTime)
	if err != nil {
		logger.Errorf("tempTime is error: %v", err)
		tempTime, _ = convert.StrTime2DateTime(`0-0-0 12:00:00`)
	}

	return &models.LoginEvent{
		EventID:       &eventId,
		EventTime:     tempTime,
		MachineUUID:   &machineUuid,
		LoginType:     &loginType,
		Username:      &username,
		SourceIP:      &ipAddress,
		SubjectUser:   &subjectUserName,
		SubjectDomain: &subjectDomain,
		ProcessName:   &processName,
	}
}

//const (
//	AccountNamePath   = "/Event/EventData/AccountName"
//	AccountDomainPath = "/Event/EventData/AccountDomain"
//	ClientNamePath    = "/Event/EventData/ClientName"
//	ClientAddressPath = "/Event/EventData/ClientAddress"
//)

func RdpEvent2RdpEventResq(event windowsLog.EventUnit) *models.RDPEventUpload {
	eventId := event[windowsLog.EventIdKey].(int64)
	eventTime := event[windowsLog.EventTimeKey].(string)
	machineUuid := event[windowsLog.MachineUUIDKey].(string)
	accountName := event[windowsLog.AccountNameKey].(string)
	accountDomain := event[windowsLog.AccountDomainKey].(string)
	clientName := event[windowsLog.ClientNameKey].(string)
	clientAddress := event[windowsLog.ClientAddressKey].(string)

	tempTime, err := convert.StrTime2DateTime(eventTime)
	if err != nil {
		logger.Errorf("tempTime is error: %v", err)
		tempTime, _ = convert.StrTime2DateTime(`0-0-0 12:00:00`)
	}

	return &models.RDPEventUpload{
		Base: &models.EventBase{
			EventID:   &eventId,
			EventTime: tempTime,
			UUID:      &machineUuid,
		},
		AccountName:   &accountName,
		AccountDomain: &accountDomain,
		ClientName:    &clientName,
		ClientAddress: &clientAddress,
	}
}

func Win32UserAccount2UserInfo(user windowsWmi.Win32_UserAccount) *models.UserInfo {
	uuid, _ := windowsWmi.QueryUuid()
	return &models.UserInfo{
		Description:  &user.Description,
		IsFocus:      nil,
		LocalAccount: &user.LocalAccount,
		Name:         &user.Name,
		Sid:          &user.SID,
		UUID:         &uuid,
	}
}
