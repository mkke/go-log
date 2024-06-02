package mlog

import (
	"os"
)

func MustSync(f *os.File) {
	_ = f.Sync()
}

type syncLogWriter struct {
}

func (writer syncLogWriter) Write(bytes []byte) (int, error) {
	// log. default device is also os.Stderr
	defer MustSync(os.Stderr)
	return os.Stderr.Write(bytes)
}
