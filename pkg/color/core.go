package color

import (
	"fmt"
	"github.com/fatih/color"
)

type config struct {
	text string
	code color.Attribute
}

var (
	ColorMap = []config{
		{text: "black", code: color.FgBlack},
		{text: "red", code: color.FgRed},
		{text: "green", code: color.FgGreen},
		{text: "yellow", code: color.FgYellow},
		{text: "blue", code: color.FgBlue},
		{text: "magent", code: color.FgMagenta},
		{text: "cyan", code: color.FgCyan},
		{text: "white", code: color.FgWhite},
		{text: "hblack", code: color.FgHiBlack},
		{text: "hred", code: color.FgHiRed},
		{text: "hgreen", code: color.FgHiGreen},
		{text: "hyellow", code: color.FgHiYellow},
		{text: "hblue", code: color.FgHiBlue},
		{text: "hmagent", code: color.FgHiMagenta},
		{text: "hcyan", code: color.FgHiCyan},
		{text: "hwhite", code: color.FgHiWhite},
	}
)

//var (
//	// 常规日志
//	ErrorPrinter = color.New(color.color.FgRed, color.Bold)
//	InfoPrinter  = color.New(color.color.FgWhite, color.Bold)
//	DebugPrinter = color.New(color.color.FgMagenta, color.Bold)
//	// 功能性日志
//	HelpPrinter   = color.New(color.color.FgCyan, color.Bold)
//	WarnPrinter   = color.New(color.color.FgYellow, color.Bold)
//	SucceedPrompt = color.New(color.color.FgGreen, color.Bold)
//)

func Red(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgRed, raw)
	fmt.Print(colored)
}

func White(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgWhite, raw)
	fmt.Print(colored)
}

func Magenta(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgMagenta, raw)
	fmt.Print(colored)
}

func Yellow(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgYellow, raw)
	fmt.Print(colored)
}

func Green(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgGreen, raw)
	fmt.Print(colored)
}

func Blue(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgBlue, raw)
	fmt.Print(colored)
}

func HRed(format string, a ...interface{}) {
	raw := fmt.Sprintf(format, a...)
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color.FgHiRed, raw)
	fmt.Print(colored)
}
