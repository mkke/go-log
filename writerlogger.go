package mlog

import (
	"fmt"
	"io"
)

type WriterLogger struct {
	io.Writer
}

func NewWriterLogger(w io.Writer) WriterLogger {
	return WriterLogger{w}
}

func (wl WriterLogger) Print(v ...interface{}) {
	_, _ = wl.Write([]byte(fmt.Sprint(v...)))
}

func (wl WriterLogger) Printf(format string, v ...interface{}) {
	_, _ = wl.Write([]byte(fmt.Sprintf(format, v...)))
}

func (wl WriterLogger) Println(v ...interface{}) {
	_, _ = wl.Write([]byte(fmt.Sprintln(v...)))
}
