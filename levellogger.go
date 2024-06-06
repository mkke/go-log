package mlog

import (
	"fmt"
	"strconv"
	"strings"
)

// Level defines log levels, corresponding to systemd/sd-daemon levels
// (see https://www.freedesktop.org/software/systemd/man/latest/sd-daemon.html).
type Level int

func (l Level) Prefix() string {
	return "<" + strconv.Itoa(int(l)) + ">"
}

func ParseLevel(s string) (Level, error) {
	switch strings.ToLower(s) {
	case "emergency":
		return LevelEmergency, nil
	case "alert":
		return LevelAlert, nil
	case "critical":
		return LevelCritical, nil
	case "error":
		return LevelError, nil
	case "warning":
		return LevelWarning, nil
	case "notice":
		return LevelNotice, nil
	case "info":
		return LevelInfo, nil
	case "debug":
		return LevelDebug, nil
	default:
		return 0, fmt.Errorf("invalid log level: %s", s)
	}
}

func (l Level) String() string {
	switch l {
	case LevelEmergency:
		return "emergency"
	case LevelAlert:
		return "alert"
	case LevelCritical:
		return "critical"
	case LevelError:
		return "error"
	case LevelWarning:
		return "warning"
	case LevelNotice:
		return "notice"
	case LevelInfo:
		return "info"
	case LevelDebug:
		return "debug"
	default:
		return l.Prefix()
	}
}

const (
	// LevelEmergency marks messages that show the system is unusable
	LevelEmergency = Level(0)

	// LevelAlert marks messages that show an action must be taken immediately
	LevelAlert = Level(1)

	// LevelCritical marks messages for critical conditions
	LevelCritical = Level(2)

	// LevelError marks messages for error conditions
	LevelError = Level(3)

	// LevelWarning marks messages for warning conditions
	LevelWarning = Level(4)

	// LevelNotice marks a normal but significant condition
	LevelNotice = Level(5)

	// LevelInfo marks informational entries
	LevelInfo = Level(6)

	// LevelDebug marks debug-level messages
	LevelDebug = Level(7)
)

// LevelLogger defines a log API including log level.
// It is meant as an optional extension for Logger implementations
type LevelLogger interface {
	// Lprint calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Lprint(level Level, v ...interface{})

	// Lprintf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Lprintf(level Level, format string, v ...interface{})

	// Lprintln calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Lprintln(level Level, v ...interface{})
}

// WithLevel wraps a Logger in a new one that logs at the given level (if logger implements LevelLogger).
// If logger doesn't implement LevelLogger, it returns the given logger unchanged.
//
// mlog.WithLevel(logger, mlog.Critical).Println(...)
// mlog.Critical(logger).Println(...)
func WithLevel(logger Logger, level Level) Logger {
	if levelLogger, ok := logger.(LevelLogger); ok {
		return LevelLoggerAdapter{
			previous: levelLogger,
			level:    level,
		}
	} else {
		return logger
	}
}

func Emergency(logger Logger) Logger {
	return WithLevel(logger, LevelEmergency)
}

func Alert(logger Logger) Logger {
	return WithLevel(logger, LevelAlert)
}

func Critical(logger Logger) Logger {
	return WithLevel(logger, LevelCritical)
}

func Error(logger Logger) Logger {
	return WithLevel(logger, LevelError)
}

func Warning(logger Logger) Logger {
	return WithLevel(logger, LevelWarning)
}

func Notice(logger Logger) Logger {
	return WithLevel(logger, LevelNotice)
}

func Info(logger Logger) Logger {
	return WithLevel(logger, LevelInfo)
}

func Debug(logger Logger) Logger {
	return WithLevel(logger, LevelDebug)
}

type LevelLoggerAdapter struct {
	previous LevelLogger
	level    Level
}

func (l LevelLoggerAdapter) Print(v ...interface{}) {
	l.previous.Lprint(l.level, v...)
}

func (l LevelLoggerAdapter) Printf(format string, v ...interface{}) {
	l.previous.Lprintf(l.level, format, v...)
}

func (l LevelLoggerAdapter) Println(v ...interface{}) {
	l.previous.Lprintln(l.level, v...)
}
