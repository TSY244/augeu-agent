package convert

import (
	"augeu-agent/pkg/swaggerCore/models"
	"augeu-agent/pkg/util/convert"
	"augeu-agent/pkg/windowsWmi"
)

type listCopyFunc[T, E any] func(T) E

func ArrayCopy[T, E any](lists []T, copyFunc listCopyFunc[T, E]) []E {
	if len(lists) == 0 {
		return nil
	}
	var res []E
	for _, item := range lists {
		res = append(res, copyFunc(item))
	}
	return res
}

// WmiPatchToMsgPatch 将 wmi 的 patch 转变成 msg patch
func WmiPatchToMsgPatch(wmiPatch windowsWmi.Win32_QuickFixEngineering) *models.Patch {
	return &models.Patch{
		Description: convert.StrPtr(wmiPatch.Description),
		HotFixID:    convert.StrPtr(wmiPatch.HotFixID),
		InstalledBy: convert.StrPtr(wmiPatch.InstalledBy),
		InstalledOn: convert.StrPtr(wmiPatch.InstalledOn),
	}
}
