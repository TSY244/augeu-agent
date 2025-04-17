package agent

import (
	"augeu-agent/internal/pkg/systeminfo"
	"augeu-agent/internal/utils/convert"
	"augeu-agent/pkg/augeuHttp"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/swaggerCore/models"
	"augeu-agent/pkg/util/utils"
	"augeu-agent/pkg/windowsLog"
	"augeu-agent/pkg/windowsWmi"
	"encoding/json"
	"github.com/0xrawsec/golang-evtx/evtx"
)

const (
	GetClientIdApiPath      = "/getClientId"
	UploadLoginEventApiPath = "/upload/loginEvent"
	UploadRdpEventApiPath   = "/upload/rdpEvent"
	UploadUsersInfoApiPath  = "/upload/usersInfo"
)

func (a *Agent) GetClientId() (string, string, error) {
	uuid, err := systeminfo.GetUuid()
	if err != nil {
		return "", "", err
	}
	ips, err := utils.GetIps()
	if err != nil {
		return "", "", err
	}
	info, err := systeminfo.GetSystemInfo()

	payload := models.GetClientIDRequest{
		Secret: &a.Conf.Secret,
		ClientInfo: &models.ClientInfo{
			UUID:       &uuid,
			IP:         *ips,
			SystemInfo: info,
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", "", err
	}

	ret, err := augeuHttp.PostRequestWithJson(a.Conf.RemoteAddr+GetClientIdApiPath, a.Header, string(jsonData))
	if err != nil {
		return "", "", err
	}
	var resp models.GetClientIDResponse
	err = json.Unmarshal([]byte(ret), &resp)
	if err != nil {
		return "", "", err
	}
	if resp == (models.GetClientIDResponse{}) {
		return "", "", err
	}
	return resp.Jwt, *resp.ClientID, nil

}

func (a *Agent) PushLoginEvent(evtxMap chan *evtx.GoEvtxMap) error {
	events := windowsLog.GetEventsForLoginEvent(evtxMap)

	resq := convert.ArrayCopy(events, convert.LoginEvent2RLoginEventResq)
	jsonData, err := json.Marshal(resq)
	if err != nil {
		logger.Errorf("PushLoginEvent json.Marshal error: %v", err)
		return err
	}
	_, err = augeuHttp.PostRequestWithJson(a.Conf.RemoteAddr+UploadLoginEventApiPath, a.Header, string(jsonData))
	if err != nil {
		logger.Errorf("PushLoginEvent PostRequestWithJson error: %v", err)
		return err
	}
	logger.Infof("PushLoginEvent success")
	return nil
}

func (a *Agent) PushRdpEvent(evtxMap chan *evtx.GoEvtxMap) error {
	events := windowsLog.GetEventsForRdpEvent(evtxMap)
	resq := convert.ArrayCopy(events, convert.RdpEvent2RdpEventResq)
	jsonData, err := json.Marshal(resq)
	if err != nil {
		logger.Errorf("PushRdpEvent json.Marshal error: %v", err)
		return err
	}
	_, err = augeuHttp.PostRequestWithJson(a.Conf.RemoteAddr+UploadRdpEventApiPath, a.Header, string(jsonData))
	if err != nil {
		logger.Errorf("PushRdpEvent PostRequestWithJson error: %v", err)
		return err
	}
	logger.Infof("PushRdpEvent success")
	return nil
}

func (a *Agent) PushUsersInfo() error {
	users, err := windowsWmi.QueryUsers()
	if err != nil {
		return err
	}
	resq := convert.ArrayCopy(users, convert.Win32UserAccount2UserInfo)
	jsonData, err := json.Marshal(resq)
	if err != nil {
		logger.Errorf("PushUsersInfo json.Marshal error: %v", err)
		return err
	}
	_, err = augeuHttp.PostRequestWithJson(a.Conf.RemoteAddr+UploadUsersInfoApiPath, a.Header, string(jsonData))
	if err != nil {
		logger.Errorf("PushUsersInfo PostRequestWithJson error: %v", err)
		return err
	}
	logger.Infof("PushUsersInfo success")
	return nil
}
