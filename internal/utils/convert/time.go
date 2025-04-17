package convert

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"time"
)

func DataTime2time(dt *strfmt.DateTime) (time.Time, error) {
	if dt == nil {
		return time.Time{}, fmt.Errorf("input is nil pointer")
	}
	// strfmt.DateTime 本质是 time.Time 的别名，可直接取值
	return time.Time(*dt), nil
}

func StrTime2time(t string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", t)
}

func StrTime2DateTime(t string) (*strfmt.DateTime, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return nil, err
	}
	// 将 time.Time 转换为 strfmt.DateTime
	dateTime := strfmt.DateTime(parsedTime)
	return &dateTime, nil
}
