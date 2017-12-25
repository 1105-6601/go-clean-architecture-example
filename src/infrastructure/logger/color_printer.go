package logger

import (
	"github.com/fatih/color"
	"time"
	"fmt"
)

type ColorPrinter struct {}

func NewColorPrinter() *ColorPrinter {
	return &ColorPrinter{}
}

func (printer *ColorPrinter) Error(text string) {
	color.Red(base(), text)
}

func (printer *ColorPrinter) ErrorF(format string, a ...interface{}) {
	color.Red(merge(format), a...)
}

func (printer *ColorPrinter) Success(text string) {
	color.Green(base(), text)
}

func (printer *ColorPrinter) SuccessF(format string, a ...interface{}) {
	color.Green(merge(format), a...)
}

func (printer *ColorPrinter) Info(text string) {
	color.Cyan(base(), text)
}

func (printer *ColorPrinter) InfoF(format string, a ...interface{}) {
	color.Cyan(merge(format), a...)
}

func (printer *ColorPrinter) Warn(text string) {
	color.Yellow(base(), text)
}

func (printer *ColorPrinter) WarnF(format string, a ...interface{}) {
	color.Yellow(merge(format), a...)
}

func (printer *ColorPrinter) Other(text string) {
	color.Magenta(base(), text)
}

func (printer *ColorPrinter) OtherF(format string, a ...interface{}) {
	color.Magenta(merge(format), a...)
}

func base() string {
	return fmt.Sprintf("[%s]", time.Now().Format("2006-01-02 15:04:05")) + " %s"
}

func merge(format string) string {
	return fmt.Sprintf(base(), format)
}
