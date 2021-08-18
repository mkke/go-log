package mlog

import (
	"github.com/fatih/color"
)

type ColoredLogger struct {
	Color    *color.Color
	Previous Logger
}

func WithColor(logColor *color.Color, previous Logger) Logger {
	return &ColoredLogger{
		Color:    logColor,
		Previous: previous,
	}
}

func (p *ColoredLogger) Print(v ...interface{}) {
	p.Previous.Print(p.Color.Sprint(v...))
}

func (p *ColoredLogger) Printf(format string, v ...interface{}) {
	p.Previous.Print(p.Color.Sprintf(format, v...))
}

func (p *ColoredLogger) Println(v ...interface{}) {
	p.Previous.Println(p.Color.Sprint(v...))
}
