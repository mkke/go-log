package mlog

import (
	"fmt"
	"os"
	"time"
)

type timestampLogWriter struct {
}

func (writer timestampLogWriter) Write(bytes []byte) (int, error) {
	// log. default device is also os.Stderr
	defer MustSync(os.Stderr)
	return fmt.Fprintf(os.Stderr, time.Now().UTC().Format("2006-01-02 15:04:05.000")+" "+string(bytes))
}
