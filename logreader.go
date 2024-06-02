package mlog

import (
	"bufio"
	"io"
	"strings"
)

// LogReader logs all lines read from a reader with an optional prefix.
//
// This is meant to be used with exec.Command:
//  if stdout, err := cmd.StdoutPipe(); err != nil {
//		return err
//	} else {
//		go LogReader(log, stdout, "stdout: ")
//	}
func LogReader(log Logger, r io.Reader, prefix ...string) {
	p := ""
	if len(prefix) == 1 {
		p = prefix[0]
	}
	s := bufio.NewScanner(r)
	for s.Scan() {
		log.Printf("%s%s\n", p, strings.TrimSpace(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Printf("%s%v\n", p, err)
	}
	if rc, ok := r.(io.ReadCloser); ok {
		_ = rc.Close()
	}
}
