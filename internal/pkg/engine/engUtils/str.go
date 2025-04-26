package engUtils

import (
	"fmt"
	"regexp"
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
	// 使用 strings.Split 分割字符串
	ret := strings.Split(str, sep)

	// 预估容量，减少切片扩容次数
	newStr := make([]string, 0, len(ret))

	for _, v := range ret {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" {
			newStr = append(newStr, trimmed)
		}
	}

	return newStr
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

func (s *StrUtils) GetPathFromCmd(cmd string) string {
	re := regexp.MustCompile(`^"([^"]+)"|^([^ "]+)`) // 匹配带引号和不带引号的路径
	matches := re.FindStringSubmatch(cmd)
	if len(matches) == 0 {
		return ""
	}

	var path string
	if matches[1] != "" { // 带引号的路径
		path = matches[1]
	} else { // 不带引号的路径
		path = matches[2]
	}
	return path
}

// GetStrSpliceLen 获取字符串切片的长度
func (s *StrUtils) GetStrSpliceLen(strSlice []string) int {
	return len(strSlice)
}

// IsStrSpliceHasSuffix 是否含有后缀
func (s *StrUtils) IsStrSpliceHasSuffix(strSlice []string, suffix string) bool {
	for _, v := range strSlice {
		if strings.HasSuffix(v, suffix) {
			return true
		}
	}
	return false
}

func (s *StrUtils) IsStrHasSuffix(strSlice string, suffix string) bool {
	return strings.HasSuffix(strSlice, suffix)
}
