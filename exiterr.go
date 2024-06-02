//go:build !tinygo

package mlog

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func ExitErr(log Logger, cmd *exec.Cmd, stdout []byte, err error) {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		prefix := ""
		if cmd.ProcessState != nil {
			if pid := cmd.ProcessState.Pid(); pid > 0 {
				prefix = "[" + strconv.Itoa(pid) + "] "
			}
		}
		log.Printf("%sexec '%s' failed: %v\n", prefix, strings.Join(cmd.Args, " "), exitErr.String())
		if len(stdout) > 0 {
			log.Printf("%sstdout: %s\n", prefix, strings.TrimSpace(string(stdout)))
		}
		if len(exitErr.Stderr) > 0 {
			log.Printf("%sstderr: %s\n", prefix, strings.TrimSpace(string(exitErr.Stderr)))
		}
	}
}
