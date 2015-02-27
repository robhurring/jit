package util

import (
	"io"
	"os"

	"github.com/wsxiaoys/terminal/color"
)

type logger struct {
	writer io.Writer
}

func (l *logger) Log(format string, a ...interface{}) {
	color.Fprintf(l.writer, format, a...)
}

func (l *logger) Write(p []byte) (n int, err error) {
	color.Fprint(l.writer, string(p))
	return
}

var (
	Logger = &logger{os.Stdout}
)
