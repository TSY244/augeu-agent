package systeminfo

import (
	"augeu-agent/internal/utils/convert"
	"augeu-agent/pkg/swaggerCore/models"
	"augeu-agent/pkg/windowsWmi"
	"runtime"
)

func GetSystemInfo() (*models.SystemInfo, error) {
	osName, err := windowsWmi.QueryOsName()
	if err != nil {
		return nil, err
	}
	osVersion, err := windowsWmi.QueryOsVersion()
	if err != nil {
		return nil, err
	}
	hotFix, err := windowsWmi.QueryHotFix()
	if err != nil {
		return nil, err
	}
	msgHotFixs := convert.ArrayCopy(hotFix, convert.WmiPatchToMsgPatch)
	return &models.SystemInfo{
		OsName:    convert.StrPtr(osName),
		OsVersion: convert.StrPtr(osVersion),
		OsArch:    convert.StrPtr(runtime.GOARCH),
		Patchs:    msgHotFixs,
	}, nil
}

func GetUuid() (string, error) {
	return windowsWmi.QueryUuid()
}
