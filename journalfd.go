package mlog

// This is copied from https://github.com/coreos/go-systemd/blob/main/journal/journal_unix.go#L102,
// which we cannot use because it does not compile for darwin.

import (
	"fmt"
	"os"
	"syscall"
)

// StderrIsJournalStream returns whether the process stderr is connected
// to the Journal's stream transport.
//
// This can be used for automatic protocol upgrading described in [Journal Native Protocol].
//
// Returns true if JOURNAL_STREAM environment variable is present,
// and stderr's device and inode numbers match it.
//
// Error is returned if unexpected error occurs: e.g. if JOURNAL_STREAM environment variable
// is present, but malformed, fstat syscall fails, etc.
//
// [Journal Native Protocol]: https://systemd.io/JOURNAL_NATIVE_PROTOCOL/#automatic-protocol-upgrading
func StderrIsJournalStream() (bool, error) {
	return fdIsJournalStream(syscall.Stderr)
}

// StdoutIsJournalStream returns whether the process stdout is connected
// to the Journal's stream transport.
//
// Returns true if JOURNAL_STREAM environment variable is present,
// and stdout's device and inode numbers match it.
//
// Error is returned if unexpected error occurs: e.g. if JOURNAL_STREAM environment variable
// is present, but malformed, fstat syscall fails, etc.
//
// Most users should probably use [StderrIsJournalStream].
func StdoutIsJournalStream() (bool, error) {
	return fdIsJournalStream(syscall.Stdout)
}

func fdIsJournalStream(fd int) (bool, error) {
	journalStream := os.Getenv("JOURNAL_STREAM")
	if journalStream == "" {
		return false, nil
	}

	var expectedStat syscall.Stat_t
	_, err := fmt.Sscanf(journalStream, "%d:%d", &expectedStat.Dev, &expectedStat.Ino)
	if err != nil {
		return false, fmt.Errorf("failed to parse JOURNAL_STREAM=%q: %v", journalStream, err)
	}

	var stat syscall.Stat_t
	err = syscall.Fstat(fd, &stat)
	if err != nil {
		return false, err
	}

	match := stat.Dev == expectedStat.Dev && stat.Ino == expectedStat.Ino
	return match, nil
}
