package mlog

import (
	"fmt"
	"log"
	"os"
)

type JournalLogger struct {
	Logger
}

var (
	_ Logger      = (*JournalLogger)(nil)
	_ LevelLogger = (*JournalLogger)(nil)
)

func NewJournalLogger() *JournalLogger {
	// journal has its own timestamps
	logger := log.New(os.Stderr, "", 0)
	logger.SetOutput(&SyncLogWriter{})
	return &JournalLogger{Logger: logger}
}

func (j JournalLogger) Lprint(level Level, v ...interface{}) {
	j.Print(append([]any{level.Prefix()}, v...))
}

func (j JournalLogger) Lprintf(level Level, format string, v ...interface{}) {
	j.Print(level.Prefix(), fmt.Sprintf(format, v...))
}

func (j JournalLogger) Lprintln(level Level, v ...interface{}) {
	j.Println(append([]any{level.Prefix()}, v...))
}
