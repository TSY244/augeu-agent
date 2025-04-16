package engUtils

import (
	"augeu-agent/pkg/color"
	"augeu-agent/pkg/logger"
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

func NewPrinter() *Printer {
	return &Printer{}
}

func (r *Printer) Info(value interface{}) {
	color.White(InfoPrint, value)
}

func (r *Printer) Warn(value interface{}) {
	color.Yellow(WarnPrint, value)
}

func (r *Printer) Debug(value interface{}) {
	color.Magenta(DebugPrinter, value)
}

func (r *Printer) Error(value interface{}) {
	color.Red(ErrorPrinter, value)
}

func (r *Printer) Remind(value interface{}) {
	color.Green(RemindPrinter, value)
}

// PrintStrSlice 打印切片
//
// 参数：
//
//	slice：切片
//	mode：打印模式，可选值为"info"、"warn"、"debug"、"error"
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
