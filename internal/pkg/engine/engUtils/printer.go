package engUtils

import (
	"augeu-agent/pkg/color"
)

var (
	basePrint    = "[task log] "
	InfoPrint    = basePrint + "[INFO]: %s\n"
	WarnPrint    = basePrint + "[WARN]: %s\n"
	DebugPrinter = basePrint + "[DEBUG]: %s\n"
	ErrorPrinter = basePrint + "[ERROR]: %s\n"
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
	color.Red(DebugPrinter, value)
}

func (r *Printer) Remind(value interface{}) {
	color.Green(DebugPrinter, value)
}

// PrintStrSlice 打印切片
//
// 参数：
//
//	slice：切片
//	mode：打印模式，可选值为"info"、"warn"、"debug"、"error"
func (r *Printer) PrintStrSlice(slice []string, mode string) {
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
	}
	for _, v := range slice {
		f(format, v)
	}
}
