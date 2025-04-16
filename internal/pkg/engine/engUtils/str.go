package engUtils

import (
	"strconv"
	"strings"
)

/*
1. 将任意类型转为str
2. 一系列的strings 方法再封装
*/

type StrUtils struct {
}

func NewStrUtils() *StrUtils {
	return &StrUtils{}
}

func (s *StrUtils) ToStr(obj interface{}) string {
	switch obj.(type) {
	case string:
		return obj.(string)
	case int:
		return strconv.Itoa(obj.(int))
	case int64:
		return strconv.FormatInt(obj.(int64), 10)
	case float64:
		return strconv.FormatFloat(obj.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(obj.(bool))
	default:
		return ""
	}
}
func (s *StrUtils) IsEmpty(str string) bool {
	return str == ""
}

func (s *StrUtils) Contains(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func (s *StrUtils) ContainsAny(str string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(str, substr) {
			return true
		}
	}
	return false
}

func (s *StrUtils) ContainsAll(str string, substrs []string) bool {
	for _, substr := range substrs {
		if !strings.Contains(str, substr) {
			return false
		}
	}
	return true
}

func (s *StrUtils) IsEuqal(str1 string, str2 string) bool {
	return str1 == str2
}

// CraterStrSlice 创建一个字符串切片
func (s *StrUtils) CraterStrSlice(strSlice ...string) []string {
	return strSlice
}

// SplitStr 切割字符串
func (s *StrUtils) SplitStr(str string, sep string) []string {
	return strings.Split(str, sep)
}

// IsStrInstrSplice 判断字符串是否在字符串切片中
func (s *StrUtils) IsStrInstrSplice(str string, strSlice []string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

// AddPrefix 给字符串添加前缀
func (s *StrUtils) AddPrefix(str string, prefix string) string {
	return prefix + str
}

// AddPrefixs 给字符串切片添加前缀
func (s *StrUtils) AddPrefixs(strSlice []string, prefix string) []string {
	for i, v := range strSlice {
		strSlice[i] = prefix + v
	}
	return strSlice
}
