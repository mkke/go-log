package mlog

import (
	"os"
)

func MustSync(f *os.File) {
	_ = f.Sync()
}

type SyncLogWriter struct {
}

func (writer SyncLogWriter) Write(bytes []byte) (int, error) {
	// log. default device is also os.Stderr
	defer MustSync(os.Stderr)
	return os.Stderr.Write(bytes)
}
