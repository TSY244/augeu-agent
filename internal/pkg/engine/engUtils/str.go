package engUtils

import (
	"fmt"
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
	return fmt.Sprintf("%v", obj)
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

// StrSliceContains 判断字符串切片中是否包含某个字符串
func (s *StrUtils) StrSliceContains(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

// GeneABackslash 生成1个反斜杠
func (s *StrUtils) GeneABackslash() string {
	return "\\"
}

// StrSliceContainsIgnoreCase 判断字符串切片中是否包含某个字符串，忽略大小写
func (s *StrUtils) StrSliceContainsIgnoreCase(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

// GetDollarSign 获取一个$符号
func (s *StrUtils) GetDollarSign() string {
	return "$"
}
