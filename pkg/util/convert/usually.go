package convert

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Str2Int(originStr string) (int, error) {
	return strconv.Atoi(originStr)
}

func Str2Bytes(originStr string) []byte {
	return []byte(originStr)
}

func Int642Str(originInt int64) string {
	return strconv.FormatInt(originInt, 10)
}

func Bytes2Str(originBytes []byte) string {
	return string(originBytes)
}

func Int32P(i int32) *int32 {
	return &i
}

func IntP(i int) *int {
	return &i
}

func Int64P(i int64) *int64 {
	return &i
}

func Int2Int64P(i int) *int64 {
	return Int64P(int64(i))
}

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func Str2Uint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func Str2Time(s string) (time.Duration, error) {
	i, err := Str2Int(s)
	if err != nil {
		return 0, err
	}
	return time.Duration(i) * time.Second, nil
}

func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func Uint642Str(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func Strings2Str(strSlice []string) string {
	return strings.Join(strSlice, ",")
}

func Any2Str(a interface{}) string {
	return fmt.Sprintf("%v", a)

}
