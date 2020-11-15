package mlog

type PrefixedLogger struct {
	Prefix   string
	Previous Logger
}

func WithPrefix(prefix string, previous Logger) Logger {
	return &PrefixedLogger{
		Prefix:   prefix,
		Previous: previous,
	}
}

func (p PrefixedLogger) Print(v ...interface{}) {
	p.Previous.Printf("%s: %v", p.Prefix, v)
}

func (p PrefixedLogger) Printf(format string, v ...interface{}) {
	p.Previous.Printf("%s: "+format, append([]interface{}{p.Prefix}, v...)...)
}

func (p PrefixedLogger) Println(v ...interface{}) {
	p.Previous.Printf("%s: %v\n", p.Prefix, v)
}
