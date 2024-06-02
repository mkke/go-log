package mlog

import (
	stdlog "log"
	"os"
)

type SyncFileLogger struct {
	logger  *stdlog.Logger
	logFile *os.File
}

func (l *SyncFileLogger) Print(v ...interface{}) {
	l.logger.Print(v...)
	_ = l.logFile.Sync()
}

func (l *SyncFileLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
	_ = l.logFile.Sync()
}

func (l *SyncFileLogger) Println(v ...interface{}) {
	l.logger.Println(v...)
	_ = l.logFile.Sync()
}

func (l *SyncFileLogger) Close() {
	_ = l.logFile.Sync()
	_ = l.logFile.Close()
}

func NewSyncFileLogger(path string) (Logger, func(), error) {
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		return nil, nil, err
	}

	newLogger := &SyncFileLogger{
		logFile: logFile,
		logger:  stdlog.New(logFile, "", stdlog.LstdFlags),
	}

	return newLogger, newLogger.Close, nil
}
