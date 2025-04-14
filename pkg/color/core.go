package color

import "github.com/fatih/color"

var (
	HelpPrinter   = color.New(color.FgCyan, color.Bold)
	WarnPrinter   = color.New(color.FgYellow, color.Bold)
	ErrorPrinter  = color.New(color.FgRed, color.Bold)
	InfoPrinter   = color.New(color.FgWhite, color.Bold)
	SucceedPrompt = color.New(color.FgGreen, color.Bold)
	DebugPrinter  = color.New(color.FgMagenta, color.Bold)
)
