package color

import "github.com/fatih/color"

var (
	// 常规日志
	ErrorPrinter = color.New(color.FgRed, color.Bold)
	InfoPrinter  = color.New(color.FgWhite, color.Bold)
	DebugPrinter = color.New(color.FgMagenta, color.Bold)
	// 功能性日志
	HelpPrinter   = color.New(color.FgCyan, color.Bold)
	WarnPrinter   = color.New(color.FgYellow, color.Bold)
	SucceedPrompt = color.New(color.FgGreen, color.Bold)
)
