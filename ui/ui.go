package ui

import (
	"io"
	"os"

	"github.com/wsxiaoys/terminal/color"
)

type UI interface {
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	Errorf(format string, a ...interface{}) (n int, err error)
	Errorln(a ...interface{}) (n int, err error)
	Write(p []byte) (n int, err error)
}

var Default UI = Console{Stdout: os.Stdout, Stderr: os.Stderr}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Default.Printf(format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return Default.Println(a...)
}

func Errorf(format string, a ...interface{}) (n int, err error) {
	return Default.Errorf(format, a...)
}

func Errorln(a ...interface{}) (n int, err error) {
	return Default.Errorln(a...)
}

func Write(p []byte) (n int, err error) {
	return Default.Write(p)
}

func Error(s string) error {
	return color.Errorf("@r%s@|", s)
}

type Console struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (c Console) Printf(format string, a ...interface{}) (n int, err error) {
	return color.Fprintf(c.Stdout, format, a...)
}

func (c Console) Println(a ...interface{}) (n int, err error) {
	return color.Fprintln(c.Stdout, a...)
}

func (c Console) Errorf(format string, a ...interface{}) (n int, err error) {
	return color.Fprintf(c.Stderr, format, a...)
}

func (c Console) Errorln(a ...interface{}) (n int, err error) {
	return color.Fprintf(c.Stderr, "@r%s@|\n", a...)
}

// Won't really work with templates since we aren't buffering at all
func (c Console) Write(p []byte) (n int, err error) {
	return c.Stdout.Write(p)
}
