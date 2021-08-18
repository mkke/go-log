package mlog

// log interface
type Logger interface {
	// Print calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Print(v ...interface{})

	// Printf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})

	// Println calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Println(v ...interface{})
}
