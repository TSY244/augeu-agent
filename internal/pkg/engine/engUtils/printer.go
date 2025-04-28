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
