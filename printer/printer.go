package printer

import (
	"os"

	"github.com/fatih/color"
)

var (
	defaultColor      = color.New(color.FgWhite, color.Reset)
	defaultInfoColor  = color.New(color.FgWhite, color.Bold)
	defaultWarnColor  = color.New(color.FgYellow, color.Bold)
	defaultErrorColor = color.New(color.FgRed, color.Bold)
)

// Printer struct
type Printer struct {
	defaultColor *color.Color
	infoColor    *color.Color
	warnColor    *color.Color
	errorColor   *color.Color
	stdout       *os.File
}

func defaultPrinter() Printer {
	printer := Printer{
		defaultColor: defaultColor,
		infoColor:    defaultInfoColor,
		warnColor:    defaultWarnColor,
		errorColor:   defaultErrorColor,
		stdout:       os.Stdout,
	}
	return printer
}

// New returns a configured printer and a close function.
// The close function deallocates an inner os.DevNull io.Writer.
//
// Usage:
//
// print, close := printer.New(false, printer.InfoColor(color.New(color.FgWhite, color.Bold)))
//
// defer close()
func New(colors ...Config) (*Printer, func()) {
	p := defaultPrinter()
	if len(colors) != 0 {
		for _, color := range colors {
			color(&p)
		}
	}
	return &p, p.Close
}

// Config can be eaily used as spread function parameter
type Config func(*Printer)

// Silent sets printer silent mode
func Silent(silent bool) Config {
	return func(p *Printer) {
		if silent {
			p.stdout, _ = os.Open(os.DevNull)
		}
	}
}

// DefaultColor sets default CLI color
func DefaultColor(c *color.Color) Config {
	return func(p *Printer) {
		p.defaultColor = c
	}
}

// InfoColor sets informational CLI color
func InfoColor(c *color.Color) Config {
	return func(p *Printer) {
		p.infoColor = c
	}
}

// WarnColor sets warn CLI color
func WarnColor(c *color.Color) Config {
	return func(p *Printer) {
		p.warnColor = c
	}
}

// ErrorColor sets error CLI color
func ErrorColor(c *color.Color) Config {
	return func(p *Printer) {
		p.errorColor = c
	}
}

// Infof formats according to a format specifier and writes to standard output using the provided info color.
func (p *Printer) Infof(format string, a ...interface{}) (n int, err error) {
	return p.infoColor.Fprintf(p.stdout, format, a...)
}

// Infoln formats using the default formats for its operands and writes to standard output using the provided info color.
func (p *Printer) Infoln(a ...interface{}) (n int, err error) {
	return p.infoColor.Fprintln(p.stdout, a...)
}

// Warnf formats according to a format specifier and writes to standard output using the provided info color.
func (p *Printer) Warnf(format string, a ...interface{}) (n int, err error) {
	return p.warnColor.Fprintf(p.stdout, format, a...)
}

// Warnln formats using the default formats for its operands and writes to standard output using the provided info color.
func (p *Printer) Warnln(a ...interface{}) (n int, err error) {
	return p.warnColor.Fprintln(p.stdout, a...)
}

// Errorf formats according to a format specifier and writes to standard output using the provided info color.
func (p *Printer) Errorf(format string, a ...interface{}) (n int, err error) {
	return p.errorColor.Fprintf(p.stdout, format, a...)
}

// Errorln formats using the default formats for its operands and writes to standard output using the provided info color.
func (p *Printer) Errorln(a ...interface{}) (n int, err error) {
	return p.errorColor.Fprintln(p.stdout, a...)
}

// Close stdout file descriptor
func (p *Printer) Close() {
	p.stdout.Close()
}

// Functions

// Infof formats according to a format specifier and writes to standard output using the provided info color.
func Infof(format string, a ...interface{}) (n int, err error) {
	return defaultInfoColor.Printf(format, a...)
}

// Infoln formats using the default formats for its operands and writes to standard output using the provided info color.
func Infoln(a ...interface{}) (n int, err error) {
	return defaultInfoColor.Println(a...)
}

// Warnf formats according to a format specifier and writes to standard output using the provided info color.
func Warnf(format string, a ...interface{}) (n int, err error) {
	return defaultWarnColor.Printf(format, a...)
}

// Warnln formats using the default formats for its operands and writes to standard output using the provided info color.
func Warnln(format string, a ...interface{}) (n int, err error) {
	return defaultWarnColor.Println(a...)
}

// Errorf formats according to a format specifier and writes to standard output using the provided info color.
func Errorf(format string, a ...interface{}) (n int, err error) {
	return defaultErrorColor.Printf(format, a...)
}

// Errorln formats using the default formats for its operands and writes to standard output using the provided info color.
func Errorln(format string, a ...interface{}) (n int, err error) {
	return defaultErrorColor.Println(a...)
}
