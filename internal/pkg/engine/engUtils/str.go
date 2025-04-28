package engUtils

import (
	"fmt"
	"regexp"
	"strings"
)

/*
1. 将任意类型转为字符串
2. 一系列的 strings 方法再封装
*/

type StrUtils struct {
}

// NewStrUtils 创建字符串工具实例
//
// return:
//   - *StrUtils 字符串工具实例
func NewStrUtils() *StrUtils {
	return &StrUtils{}
}

// ToStr 将任意类型转换为字符串
//
// params:
//   - obj 任意类型的对象
//
// return:
//   - string 转换后的字符串
func (s *StrUtils) ToStr(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

// IsEmpty 判断字符串是否为空
//
// params:
//   - str 字符串
//
// return:
//   - bool 如果字符串为空，返回 true；否则返回 false
func (s *StrUtils) IsEmpty(str string) bool {
	return str == ""
}

// Contains 判断字符串是否包含子字符串
//
// params:
//   - str 字符串
//   - substr 子字符串
//
// return:
//   - bool 如果字符串包含子字符串，返回 true；否则返回 false
func (s *StrUtils) Contains(str string, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsAny 判断字符串是否包含任意一个子字符串
//
// params:
//   - str 字符串
//   - substrs 子字符串切片
//
// return:
//   - bool 如果字符串包含任意一个子字符串，返回 true；否则返回 false
func (s *StrUtils) ContainsAny(str string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(str, substr) {
			return true
		}
	}
	return false
}

// ContainsAll 判断字符串是否包含所有子字符串
//
// params:
//   - str 字符串
//   - substrs 子字符串切片
//
// return:
//   - bool 如果字符串包含所有子字符串，返回 true；否则返回 false
func (s *StrUtils) ContainsAll(str string, substrs []string) bool {
	for _, substr := range substrs {
		if !strings.Contains(str, substr) {
			return false
		}
	}
	return true
}

// IsEuqal 判断两个字符串是否相等
//
// params:
//   - str1 字符串1
//   - str2 字符串2
//
// return:
//   - bool 如果两个字符串相等，返回 true；否则返回 false
func (s *StrUtils) IsEuqal(str1 string, str2 string) bool {
	return str1 == str2
}

// CraterStrSlice 创建一个字符串切片
//
// params:
//   - strSlice 可变参数，表示字符串切片的元素
//
// return:
//   - []string 返回创建的字符串切片
func (s *StrUtils) CraterStrSlice(strSlice ...string) []string {
	return strSlice
}

// SplitStr 切割字符串
//
// params:
//   - str 待切割的字符串
//   - sep 分隔符
//
// return:
//   - []string 返回根据分隔符切割后的字符串切片（去除空格和空元素）
func (s *StrUtils) SplitStr(str string, sep string) []string {
	ret := strings.Split(str, sep)
	newStr := make([]string, 0, len(ret))
	for _, v := range ret {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" {
			newStr = append(newStr, trimmed)
		}
	}
	return newStr
}

// IsStrinstrSplice 判断字符串是否在字符串切片中
//
// params:
//   - str 字符串
//   - strSlice 字符串切片
//
// return:
//   - bool 如果字符串存在于字符串切片中，返回 true；否则返回 false
func (s *StrUtils) IsStrinstrSplice(str string, strSlice []string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

// AddPrefix 给字符串添加前缀
//
// params:
//   - str 字符串
//   - prefix 前缀
//
// return:
//   - string 添加前缀后的字符串
func (s *StrUtils) AddPrefix(str string, prefix string) string {
	return prefix + str
}

// AddPrefixs 给字符串切片中的每个字符串添加前缀
//
// params:
//   - strSlice 字符串切片
//   - prefix 前缀
//
// return:
//   - []string 添加前缀后的字符串切片
func (s *StrUtils) AddPrefixs(strSlice []string, prefix string) []string {
	for i, v := range strSlice {
		strSlice[i] = prefix + v
	}
	return strSlice
}

// StrSliceContains 判断字符串切片中是否包含某个字符串
//
// params:
//   - strSlice 字符串切片
//   - str 字符串
//
// return:
//   - bool 如果字符串存在于字符串切片中，返回 true；否则返回 false
func (s *StrUtils) StrSliceContains(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

// GeneABackslash 生成一个反斜杠
//
// return:
//   - string 返回反斜杠字符串
func (s *StrUtils) GeneABackslash() string {
	return "\\"
}

// StrSliceContainsIgnoreCase 判断字符串切片中是否包含某个字符串，忽略大小写
//
// params:
//   - strSlice 字符串切片
//   - str 字符串
//
// return:
//   - bool 如果字符串存在于字符串切片中（忽略大小写），返回 true；否则返回 false
func (s *StrUtils) StrSliceContainsIgnoreCase(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

// GetDollarSign 获取一个 $ 符号
//
// return:
//   - string 返回 $ 符号字符串
func (s *StrUtils) GetDollarSign() string {
	return "$"
}

// GetPathFromCmd 从命令字符串中提取路径
//
// params:
//   - cmd 命令字符串（如 "\"C:\\Program Files\\app.exe\" arg1"）
//
// return:
//   - string 提取的路径（如 "C:\\Program Files\\app.exe"），未匹配时返回空字符串
//
// notice:
//  1. 支持带引号和不带引号的路径格式
//  2. 使用正则表达式匹配路径部分
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
//
// params:
//   - strSlice 字符串切片
//
// return:
//   - int 字符串切片的长度
func (s *StrUtils) GetStrSpliceLen(strSlice []string) int {
	return len(strSlice)
}

// IsStrSpliceHasSuffix 判断字符串切片中是否有以指定后缀结尾的字符串
//
// params:
//   - strSlice 字符串切片
//   - suffix 后缀
//
// return:
//   - bool 如果字符串切片中有以指定后缀结尾的字符串，返回 true；否则返回 false
func (s *StrUtils) IsStrSpliceHasSuffix(strSlice []string, suffix string) bool {
	for _, v := range strSlice {
		if strings.HasSuffix(v, suffix) {
			return true
		}
	}
	return false
}

// IsStrHasSuffix 判断字符串是否以指定后缀结尾
//
// params:
//   - str 字符串
//   - suffix 后缀
//
// return:
//   - bool 如果字符串以指定后缀结尾，返回 true；否则返回 false
func (s *StrUtils) IsStrHasSuffix(strSlice string, suffix string) bool {
	return strings.HasSuffix(strSlice, suffix)
}
