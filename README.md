# 使用



```
.\AugetAgent.exe -h
Usage of AugetAgent.exe:
  -cp string
        配置文件路径
  -mode string
        加载方式: basic, remote, local
                basic: 从程序内部加载默认规则
                remote: 加载远端提供的配置文件, -cp 参数需要提供一个远程url
                local: 从本地加载一个配置文件, -cp 需要提供一个本地文件路径
                 (default "basic")
  -r string
        server api 地址，请注意加上/api/v1，举个栗子：http://127.0.0.1/api/v1
  -s string
        server 提供的 secret
  -wbapi string
        微步api密钥
  -ws string
        server 提供的websocket 地址，注意使用ws://开头
```

## 本地模式

本地模式启动

本地模式使用的是程序内置的规则

直接运行即可

```
.\AugetAgent.exe -wbapi  XXXXXXXXXXXXXX
```

或者是使用`-mode basic` 参数进行指定

![image-20250428221210288](https://raw.githubusercontent.com/TSY244/augeu-agent/refs/heads/master/img/image-20250428221059918.png)







## 远程文件模式

使用python 起一个服务，把文件挂载，然后使用`-mode remoteFile` 的模式进行加载

```
.\AugetAgent.exe -mode remoteFile -cp http://127.0.0.1:8000/rule.txt -wbapi  XXXXXXXXXXXXXX
```



## 远程api 模式

先启动backend

![image-20250428224909131.png (1633×68)](https://raw.githubusercontent.com/TSY244/augeu-agent/refs/heads/master/img/image-20250428224909131.png)

这个token 就是secrete

```
.\AugetAgent.exe -mode remoteApi -r http://127.0.0.1:8080/api/v1 -s R1gmbk9yPiNEY3RYNVdSd -wbapi  XXXXXXXXXXXXXX
```

## 监听模式

``` 
.\AugeuAgent.exe -mode monitor -target ip/domain
```



![image-20250504105442370](https://raw.githubusercontent.com/TSY244/augeu-agent/refs/heads/master/img/image-20250504105442370.png)

# 日志输出分析

## rule 内部日志

### 标记

每一个rule 内部输出的日志

> [task log] [REMIND] [自启动文件夹下检测1] desktop.ini

带有前缀 `[task log] `

![image-20250428224944255](https://raw.githubusercontent.com/TSY244/augeu-agent/refs/heads/master/img/image-20250428224944255.png)

## 程序日志

### 格式

采用格式

```
时间戳 level error
2025-04-28T22:47:48+08:00 ERR get hash error: open explorer.exe: The system cannot find the file specified.
```

![image-20250428224909131](https://raw.githubusercontent.com/TSY244/augeu-agent/refs/heads/master/img/image-20250428224909131.png)

### 常见日志解析

1. reg err

   ``` 
   2025-04-28T22:47:48+08:00 ERR get hash error: open explorer.exe: The system cannot find the file specified.
   ```

   这个日志输出表示没有这个注册表路径，有可能是错误。有可能是真没有危害



# 扩展

## 代码上扩展

### 添加模块

rule 中的代码主要是使用模板和反射进行加载

所以只支持写好的规则，然后进行注册，最后在rule 中使用

internal/pkg/agent/core.go

```go

func NewAgent(c *param.Config) *Agent {
	checkConf(c)
	rootCtx, cancel := context.WithCancel(context.Background())
	
	agent := &Agent{
		Conf:    c,
		RootCtx: rootCtx,
		Cancel:  cancel,
		//Eng:  engine.NewEngine(),
	}
	agent.ApiOuter = map[string]interface{}{
		"println":      fmt.Println,
		"reg":          engUtils.NewReg(),
		"strUtils":     engUtils.NewStrUtils(),
		"fileSysUtils": engUtils.NewFileSys(),
		"printer":      engUtils.NewPrinter(),
		"base":         engUtils.NewBase(),
		"agent":        agent,
		"weibu":        engUtils.NewWeiBuUtils(),
		"check":        engUtils.NewCheck(),
		"service":      engUtils.NewService(),
		"err":          consts.ErrKey,
		"schedule":     engUtils.NewSchedule(),
		"phs":          engUtils.NewPhs(),
	}
	return agent
}
```

在`agent.ApiOuter` 这个map 中进行注册

key 为rule 中调用的时候的名字

value 是函数指针，或者是结构体指针。建议放结构体，然后函数是成员函数的形式。

举个例子，比如reg

internal/pkg/engine/engUtils/reg.go

```go
package engUtils

import (
	"augeu-agent/internal/utils/consts"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/registration"
	"golang.org/x/sys/windows/registry"
	"regexp"
)

type Reg struct {
}

func NewReg() *Reg {
	return &Reg{}
}

func (r *Reg) GetPathSubKeys(path string) []string {
	names, err := registration.GetPathSubKeys(path)
	if err != nil {
		logger.Errorf("get path sub keys error: %v", err)
		return nil
	}
	return names
}

func (r *Reg) GetRegPathValueNames(path string) []string {
	//return registration.GetRegPathValueNames(path)
	names, err := registration.GetRegPathValueNames(path)
	if err != nil {
		logger.Errorf("get reg path value names error: %v", err)
		return nil
	}
	return names
}

func (r *Reg) GetRegPathValue(path string, name string) string {
	//return registration.GetRegPathStringValue(path, name)
	value, err := registration.GetRegPathStringValue(path, name)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return value
}

func (r *Reg) RegNameType(reg registry.Key, name string) uint32 {
	//return registration.RegNameType(reg, name)
	ret, err := registration.RegNameType(reg, name)
	if err != nil {
		logger.Errorf("get reg name type error: %v", err)
		return 0
	}
	return ret
}

func (r *Reg) IsPathWithSubKey(path string, subKey string) bool {
	return registration.IsPathWithSubKey(path, subKey)
}
func (r *Reg) IsPathWithName(path string, name string) bool {
	return registration.IsPathWithName(path, name)
}

func (r *Reg) GetPathFromCmd(cmd string) string {
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

func (r *Reg) IsHavePath(path string) bool {
	return registration.IsHavePath(path)
}

func (r *Reg) GetDefaultRegPathValue(path string) string {
	ret, err := registration.GetDefaultRegPathValue(path)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return ret
}

```



### 模块介绍

#### check

```go 
package engUtils

import (
	"augeu-agent/internal/pkg/weibu"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/util/convert"
	"fmt"
	"strings"
)

type Check struct {
}

// NewCheck 创建检查工具实例
//
// return:
//   - *Check 检查工具实例
func NewCheck() *Check {
	return &Check{}
}

// CheckHash 通过微步检查哈希值的风险率是否低于指定阈值
//
// params:
//   - hash 哈希值（如文件的MD5、SHA256等）
//   - a 配置对象，包含多引擎检测相关配置
//   - proxy 代理地址（可选，用于网络请求）
//   - rate 风险率阈值（如0.5表示50%）
//
// return:
//   - bool 是否通过检查（风险率 <= rate为true，否则为false）
//
// notice:
//  1. 如果获取多引擎检测结果失败，返回false
//  2. 如果返回结果格式不正确（如不包含"/"或分割后长度不为2），返回false
//  3. 如果字符串转整数失败，返回false
//  4. 风险率计算公式：intRate1 / intRate2，其中intRate1和intRate2为分割后的两个整数
func (c *Check) CheckHash(hash string, a *weibu.Config, proxy string, rate float64) bool {
	ret, err := weibu.GetMultiEngines(hash, a, proxy)
	if err != nil {
		logger.Errorf("CheckHash error is %v", err)
		return false
	}
	if ret == "" {
		return false
	}
	if !strings.Contains(ret, "/") {
		logger.Errorf("CheckHash error is %v", fmt.Errorf("ret is %s", ret))
		return false
	}
	rates := strings.Split(ret, "/")
	if len(rates) != 2 {
		logger.Errorf("CheckHash error is %v", fmt.Errorf("ret is %s", ret))
		return false
	}
	intRate1, err := convert.Str2Int(rates[0])
	if err != nil {
		logger.Errorf("CheckHash error is %v", err)
		return false
	}
	intRate2, err := convert.Str2Int(rates[1])
	if err != nil {
		logger.Errorf("CheckHash error is %v", err)
		return false
	}
	return float64(intRate1)/float64(intRate2) <= rate
}

```

#### reg

``` go
package engUtils

import (
	"augeu-agent/internal/utils/consts"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/registration"
	"golang.org/x/sys/windows/registry"
	"regexp"
)

type Reg struct {
}

// NewReg 创建注册表操作对象
//
// return:
//   - *Reg 注册表操作实例
func NewReg() *Reg {
	return &Reg{}
}

// GetRegPathValueNames 获取注册表路径下的所有值名称
//
// params:
//   - path 注册表路径
//
// return:
//   - []string 值名称列表
//
// notice:
//  1. 发生错误时返回nil
func (r *Reg) GetRegPathValueNames(path string) []string {
	names, err := registration.GetRegPathValueNames(path)
	if err != nil {
		logger.Errorf("get reg path value names error: %v", err)
		return nil
	}
	return names
}

// GetRegPathValue 获取注册表路径指定值的字符串值
//
// params:
//   - path 注册表路径
//   - name 值名称
//
// return:
//   - string 值内容，错误时返回consts.ErrKey
//
// notice:
//  1. 非字符串类型值将返回错误
//  2. 错误时返回预定义错误标识consts.ErrKey
func (r *Reg) GetRegPathValue(path string, name string) string {
	value, err := registration.GetRegPathStringValue(path, name)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return value
}

// RegNameType 获取注册表键中指定值的类型
//
// params:
//   - reg 注册表键对象
//   - name 值名称
//
// return:
//   - uint32 注册表值类型代码（如registry.REG_SZ等）
//
// notice:
//  1. 错误时返回0
func (r *Reg) RegNameType(reg registry.Key, name string) uint32 {
	ret, err := registration.RegNameType(reg, name)
	if err != nil {
		logger.Errorf("get reg name type error: %v", err)
		return 0
	}
	return ret
}

// IsPathWithSubKey 检查注册表路径是否存在指定子键
//
// params:
//   - path 父路径
//   - subKey 子键名称
//
// return:
//   - bool 是否存在
func (r *Reg) IsPathWithSubKey(path string, subKey string) bool {
	return registration.IsPathWithSubKey(path, subKey)
}

// IsPathWithName 检查注册表路径是否存在指定值名称
//
// params:
//   - path 注册表路径
//   - name 值名称
//
// return:
//   - bool 是否存在
func (r *Reg) IsPathWithName(path string, name string) bool {
	return registration.IsPathWithName(path, name)
}

// GetPathFromCmd 从命令字符串中提取路径
//
// params:
//   - cmd 命令字符串（如"\"C:\Program Files\app.exe\" arg1\"）
//
// return:
//   - string 提取的路径（如"C:\Program Files\app.exe"）
//
// notice:
//  1. 支持带引号和不带引号的路径格式
//  2. 未匹配时返回空字符串
func (r *Reg) GetPathFromCmd(cmd string) string {
	re := regexp.MustCompile(`^"([^"]+)"|^([^ "]+)`)
	matches := re.FindStringSubmatch(cmd)
	if len(matches) == 0 {
		return ""
	}
	if matches[1] != "" {
		return matches[1]
	}
	return matches[2]
}

// IsHavePath 检查注册表路径是否存在
//
// params:
//   - path 注册表路径
//
// return:
//   - bool 路径是否存在
func (r *Reg) IsHavePath(path string) bool {
	return registration.IsHavePath(path)
}

// GetDefaultRegPathValue 获取注册表路径的默认值
//
// params:
//   - path 注册表路径
//
// return:
//   - string 默认值内容，错误时返回consts.ErrKey
//
// notice:
//  1. 错误时返回预定义错误标识consts.ErrKey
func (r *Reg) GetDefaultRegPathValue(path string) string {
	ret, err := registration.GetDefaultRegPathValue(path)
	if err != nil {
		logger.Errorf("get reg path value error: %v", err)
		return consts.ErrKey
	}
	return ret
}

```

#### str

``` go
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

```

#### strUtils

```go 
package engUtils

import (
	"augeu-agent/pkg/fileSystem"
	"augeu-agent/pkg/logger"
	"os"
)

type FileSysUtils struct {
}

// NewFileSys 创建文件系统工具实例
//
// return:
//   - *FileSysUtils 文件系统工具实例
func NewFileSys() *FileSysUtils {
	return &FileSysUtils{}
}

// LsFile 获取指定路径下的所有文件名
//
// params:
//   - path 路径
//
// return:
//   - []string 文件名列表，错误时返回nil
//
// notice:
//  1. 如果路径无效或发生错误，返回nil
func (f *FileSysUtils) LsFile(path string) []string {
	files, err := fileSystem.LsFile(path)
	if err != nil {
		logger.Errorf("get path file error: %v", err)
		return nil
	}
	return files
}

// LsDir 获取指定路径下的所有目录名
//
// params:
//   - path 路径
//
// return:
//   - []string 目录名列表，错误时返回nil
//
// notice:
//  1. 如果路径无效或发生错误，返回nil
func (f *FileSysUtils) LsDir(path string) []string {
	dirs, err := fileSystem.LsDir(path)
	if err != nil {
		logger.Errorf("get path dir error: %v", err)
		return nil
	}
	return dirs
}

// GetHashWithFilePath 获取指定文件的哈希值
//
// params:
//   - filePath 文件路径
//
// return:
//   - string 哈希值，错误时返回空字符串
//
// notice:
//  1. 如果文件不存在或计算哈希失败，返回空字符串
func (f *FileSysUtils) GetHashWithFilePath(filePath string) string {
	hash, err := fileSystem.GetHashWithFilePath(filePath)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return ""
	}
	return hash
}

// GetHashesWithFilePaths 批量获取文件哈希值
//
// params:
//   - filePaths 文件路径列表
//
// return:
//   - []string 哈希值列表，错误时返回nil
//
// notice:
//  1. 如果任意文件哈希计算失败，整体返回nil
func (f *FileSysUtils) GetHashesWithFilePaths(filePaths []string) []string {
	hashes, err := fileSystem.GetHashsWithFilePaths(filePaths)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return nil
	}
	return hashes
}

// IntoFile 创建文件并写入数据
//
// params:
//   - fileName 文件名
//   - data 写入的数据
//
// return:
//   - error 错误信息，成功时为nil
//
// notice:
//  1. 如果文件创建失败或写入失败，返回相应错误
func (f *FileSysUtils) IntoFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		logger.Errorf("write file error: %v", err)
		return err
	}
	return nil
}

// AddToFile 向文件追加数据
//
// params:
//   - fileName 文件名
//   - data 追加的数据
//
// return:
//   - error 错误信息，成功时为nil
//
// notice:
//  1. 如果文件打开失败或追加数据失败，返回相应错误
func (f *FileSysUtils) AddToFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		logger.Errorf("write file error: %v", err)
		return err
	}
	return nil
}

// StrSpliceIntoFile 将字符串切片逐行写入文件
//
// params:
//   - fileName 文件名
//   - strSlice 字符串切片
//
// notice:
//  1. 如果文件创建失败或写入失败，记录错误日志
//  2. 每个字符串会以换行符分隔写入文件
func (f *FileSysUtils) StrSpliceIntoFile(fileName string, strSlice []string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return
	}
	defer file.Close()
	for _, str := range strSlice {
		_, err = file.WriteString(str + "\n")
		if err != nil {
			logger.Errorf("write file error: %v", err)
		}
	}
}

```



#### fileSysUtils

```go
package engUtils

import (
	"augeu-agent/pkg/fileSystem"
	"augeu-agent/pkg/logger"
	"os"
)

type FileSysUtils struct {
}

// NewFileSys 创建文件系统工具实例
//
// return:
//   - *FileSysUtils 文件系统工具实例
func NewFileSys() *FileSysUtils {
	return &FileSysUtils{}
}

// LsFile 获取指定路径下的所有文件名
//
// params:
//   - path 路径
//
// return:
//   - []string 文件名列表，错误时返回nil
//
// notice:
//  1. 如果路径无效或发生错误，返回nil
func (f *FileSysUtils) LsFile(path string) []string {
	files, err := fileSystem.LsFile(path)
	if err != nil {
		logger.Errorf("get path file error: %v", err)
		return nil
	}
	return files
}

// LsDir 获取指定路径下的所有目录名
//
// params:
//   - path 路径
//
// return:
//   - []string 目录名列表，错误时返回nil
//
// notice:
//  1. 如果路径无效或发生错误，返回nil
func (f *FileSysUtils) LsDir(path string) []string {
	dirs, err := fileSystem.LsDir(path)
	if err != nil {
		logger.Errorf("get path dir error: %v", err)
		return nil
	}
	return dirs
}

// GetHashWithFilePath 获取指定文件的哈希值
//
// params:
//   - filePath 文件路径
//
// return:
//   - string 哈希值，错误时返回空字符串
//
// notice:
//  1. 如果文件不存在或计算哈希失败，返回空字符串
func (f *FileSysUtils) GetHashWithFilePath(filePath string) string {
	hash, err := fileSystem.GetHashWithFilePath(filePath)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return ""
	}
	return hash
}

// GetHashesWithFilePaths 批量获取文件哈希值
//
// params:
//   - filePaths 文件路径列表
//
// return:
//   - []string 哈希值列表，错误时返回nil
//
// notice:
//  1. 如果任意文件哈希计算失败，整体返回nil
func (f *FileSysUtils) GetHashesWithFilePaths(filePaths []string) []string {
	hashes, err := fileSystem.GetHashsWithFilePaths(filePaths)
	if err != nil {
		logger.Errorf("get hash error: %v", err)
		return nil
	}
	return hashes
}

// IntoFile 创建文件并写入数据
//
// params:
//   - fileName 文件名
//   - data 写入的数据
//
// return:
//   - error 错误信息，成功时为nil
//
// notice:
//  1. 如果文件创建失败或写入失败，返回相应错误
func (f *FileSysUtils) IntoFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		logger.Errorf("write file error: %v", err)
		return err
	}
	return nil
}

// AddToFile 向文件追加数据
//
// params:
//   - fileName 文件名
//   - data 追加的数据
//
// return:
//   - error 错误信息，成功时为nil
//
// notice:
//  1. 如果文件打开失败或追加数据失败，返回相应错误
func (f *FileSysUtils) AddToFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		logger.Errorf("write file error: %v", err)
		return err
	}
	return nil
}

// StrSpliceIntoFile 将字符串切片逐行写入文件
//
// params:
//   - fileName 文件名
//   - strSlice 字符串切片
//
// notice:
//  1. 如果文件创建失败或写入失败，记录错误日志
//  2. 每个字符串会以换行符分隔写入文件
func (f *FileSysUtils) StrSpliceIntoFile(fileName string, strSlice []string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("create file error: %v", err)
		return
	}
	defer file.Close()
	for _, str := range strSlice {
		_, err = file.WriteString(str + "\n")
		if err != nil {
			logger.Errorf("write file error: %v", err)
		}
	}
}

```

#### phs

```go 
package engUtils

import (
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/powershell"
	"fmt"
)

type Phs struct {
}

// NewPhs 创建 PowerShell 工具实例
//
// return:
//   - *Phs PowerShell 工具实例
func NewPhs() *Phs {
	return &Phs{}
}

// GetScheduledTaskCommands 获取计划任务的命令信息
//
// return:
//   - []string 包含计划任务命令信息的字符串列表，错误时返回nil
//
// notice:
//  1. 使用 powershell.GetBitsAdminInfo 获取计划任务信息
//  2. 如果获取信息失败，记录错误日志并返回nil
//  3. 每个计划任务的信息格式为："id: <JobId> 任务名: <DisplayName> 传输类型: <TransferType> 任务状态: <JobState>"
func (p *Phs) GetScheduledTaskCommands() []string {
	ret, err := powershell.GetBitsAdminInfo()
	if err != nil {
		logger.Errorf("获取计划任务命令失败: %v", err)
		return nil
	} else {
		var retStr []string
		for _, v := range ret {
			strValue := fmt.Sprintf("id: %s 任务名: %s 传输类型: %s 任务状态: %s ",
				v.JobId, v.DisplayName, v.TransferType, v.JobState)
			retStr = append(retStr, strValue)
		}
		return retStr
	}
}

```



#### printer

```go 
package engUtils

import (
	"augeu-agent/pkg/color"
	"augeu-agent/pkg/logger"
	"augeu-agent/pkg/util/convert"
	"strings"
)

var (
	basePrint     = "[task log] "
	InfoPrint     = basePrint + "[INFO] %s\n"
	WarnPrint     = basePrint + "[WARN] %s\n"
	DebugPrinter  = basePrint + "[DEBUG] %s\n"
	ErrorPrinter  = basePrint + "[ERROR] %s\n"
	RemindPrinter = basePrint + "[REMIND] %s\n"
)

type Printer struct {
}

// NewPrinter 创建打印工具实例
//
// return:
//   - *Printer 打印工具实例
func NewPrinter() *Printer {
	return &Printer{}
}

// Info 打印信息日志
//
// params:
//   - value 日志内容（可以是任意类型）
//   - addInfo 可选的附加信息（多个字符串组成的切片）
//
// notice:
//  1. 如果提供了附加信息，会将附加信息以"[info1,info2]"的形式添加到日志前缀
//  2. 使用白色字体打印日志
func (r *Printer) Info(value interface{}, addInfo ...string) {
	if len(addInfo) > 0 {
		value = "[" + strings.Join(addInfo, ",") + "] " + convert.Any2Str(value)
	}
	raw := convert.Any2Str(value)
	color.White(InfoPrint, raw)
}

// Warn 打印警告日志
//
// params:
//   - value 日志内容（可以是任意类型）
//   - addInfo 可选的附加信息（多个字符串组成的切片）
//
// notice:
//  1. 如果提供了附加信息，会将附加信息以"[info1,info2]"的形式添加到日志前缀
//  2. 使用黄色字体打印日志
func (r *Printer) Warn(value interface{}, addInfo ...string) {
	if len(addInfo) > 0 {
		value = "[" + strings.Join(addInfo, ",") + "] " + convert.Any2Str(value)
	}
	raw := convert.Any2Str(value)
	color.Yellow(WarnPrint, raw)
}

// Debug 打印调试日志
//
// params:
//   - value 日志内容（可以是任意类型）
//   - addInfo 可选的附加信息（多个字符串组成的切片）
//
// notice:
//  1. 如果提供了附加信息，会将附加信息以"[info1,info2]"的形式添加到日志前缀
//  2. 使用洋红色字体打印日志
func (r *Printer) Debug(value interface{}, addInfo ...string) {
	if len(addInfo) > 0 {
		value = "[" + strings.Join(addInfo, ",") + "] " + convert.Any2Str(value)
	}
	raw := convert.Any2Str(value)
	color.Magenta(DebugPrinter, raw)
}

// Error 打印错误日志
//
// params:
//   - value 日志内容（可以是任意类型）
//   - addInfo 可选的附加信息（多个字符串组成的切片）
//
// notice:
//  1. 如果提供了附加信息，会将附加信息以"[info1,info2]"的形式添加到日志前缀
//  2. 使用红色字体打印日志
func (r *Printer) Error(value interface{}, addInfo ...string) {
	if len(addInfo) > 0 {
		value = "[" + strings.Join(addInfo, ",") + "] " + convert.Any2Str(value)
	}
	raw := convert.Any2Str(value)
	color.Red(ErrorPrinter, raw)
}

// Remind 打印提醒日志
//
// params:
//   - value 日志内容（可以是任意类型）
//   - addInfo 可选的附加信息（多个字符串组成的切片）
//
// notice:
//  1. 如果提供了附加信息，会将附加信息以"[info1,info2]"的形式添加到日志前缀
//  2. 使用绿色字体打印日志
func (r *Printer) Remind(value interface{}, addInfo ...string) {
	if len(addInfo) > 0 {
		value = "[" + strings.Join(addInfo, ",") + "] " + convert.Any2Str(value)
	}
	raw := convert.Any2Str(value)
	color.Green(RemindPrinter, raw)
}

// PrintStrSlice 打印字符串切片
//
// params:
//   - slice 字符串切片
//   - mode 打印模式，可选值为"info"、"warn"、"debug"、"error"、"remind"
//   - addInfo 可选的附加信息（多个字符串组成的切片）
//
// notice:
//  1. 根据mode选择不同的打印颜色和格式
//  2. 如果mode无效，记录错误日志并提示有效模式
//  3. 如果提供了附加信息，会将附加信息以"[info1,info2]"的形式添加到每条日志前缀
func (r *Printer) PrintStrSlice(slice []string, mode string, addInfo ...string) {
	addInfoStr := ""
	if len(addInfo) > 0 {
		addInfoStr = "[" + strings.Join(addInfo, ",") + "] "
	}
	var f func(format string, a ...interface{})
	var format string
	switch mode {
	case "info":
		f = color.White
		format = InfoPrint
	case "warn":
		f = color.Yellow
		format = WarnPrint
	case "debug":
		f = color.Magenta
		format = DebugPrinter
	case "error":
		f = color.Red
		format = ErrorPrinter
	case "remind":
		f = color.Green
		format = RemindPrinter
	default:
		logger.Errorf("unknown mode: %s", mode)
		logger.Infof("modes: info, warn, debug, error, remind")
		return
	}
	for _, v := range slice {
		f(format, addInfoStr+v)
	}
}

```

#### base

```go
package engUtils

type Base struct {
}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) SizeForStr(str []string) int {
	return len(str)
}

func (b *Base) GeneFileSegmentation(chunkSize int, src string) string {
	var dst string
	for i := 0; i < chunkSize; i++ {
		dst += src
	}
	return dst
}
```



#### weibu

```go
package engUtils

import (
	"augeu-agent/internal/pkg/weibu"
)

type WeiBuUtils struct {
}

// NewWeiBuUtils 创建 WeiBu 工具实例
//
// return:
//   - *WeiBuUtils WeiBu 工具实例
func NewWeiBuUtils() *WeiBuUtils {
	return &WeiBuUtils{}
}

// GetFileReport 获取单个文件的报告
//
// params:
//   - target 目标文件路径或标识符
//   - a 配置对象，包含与 WeiBu 服务交互所需的参数
//   - proxy 代理地址（可选）
//
// return:
//   - string 文件报告的结果
//   - error 如果获取报告失败，返回错误信息；否则返回 nil
//
// notice:
//  1. 调用 weibu.GetFileReport 方法获取单个文件的报告
//  2. 如果发生错误，返回错误信息以便调用方处理
func (wb *WeiBuUtils) GetFileReport(target string, a *weibu.Config, proxy string) (string, error) {
	return weibu.GetFileReport(target, a, proxy)
}

// GetFilesReport 获取多个文件的报告
//
// params:
//   - targets 目标文件路径或标识符的切片
//   - a 配置对象，包含与 WeiBu 服务交互所需的参数
//   - proxy 可选的代理地址（支持多个代理地址）
//
// return:
//   - []string 文件报告结果的切片
//   - error 如果获取报告失败，返回错误信息；否则返回 nil
//
// notice:
//  1. 调用 weibu.GetFilesReport 方法获取多个文件的报告
//  2. 支持传入多个代理地址（通过可变参数 proxy 实现）
//  3. 如果发生错误，返回错误信息以便调用方处理
func (wb *WeiBuUtils) GetFilesReport(targets []string, a *weibu.Config, proxy ...string) ([]string, error) {
	return weibu.GetFilesReport(targets, a, proxy...)
}

```



## 规则扩展

主要就是使用上面的提到的模块重新组合，产生新的规则

