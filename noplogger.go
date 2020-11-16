package mlog

type NopLogger struct {
}

func NewNopLogger() Logger {
	return &NopLogger{}
}

func (n *NopLogger) Print(v ...interface{}) {
}

func (n *NopLogger) Printf(format string, v ...interface{}) {
}

func (n *NopLogger) Println(v ...interface{}) {
}
