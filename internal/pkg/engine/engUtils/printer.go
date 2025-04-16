package engUtils

import "augeu-agent/pkg/color"

var (
	InfoPrint    = "In task [INFO]: %s\n"
	WarnPrint    = "In task [WARN]: %s\n"
	DebugPrinter = "In task [DEBUG]: %s\n"
)

type Printer struct {
}

func NewPrinter() *Printer {
	return &Printer{}
}

func (r *Printer) Info(value interface{}) {
	color.InfoPrinter.Printf(InfoPrint, value)
}

func (r *Printer) Warn(value interface{}) {
	color.WarnPrinter.Printf(WarnPrint, value)
}

func (r *Printer) Debug(value interface{}) {
	color.DebugPrinter.Printf(DebugPrinter, value)
}
