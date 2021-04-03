package logger

import (
	lg "log"
	"os"
)

var (
	// Info - log for common info
	Info = lg.New(os.Stdout, "[INFO] ", lg.Lshortfile)
	// Error - log for errors
	Error = lg.New(os.Stderr, "[ERROR] ", lg.Lshortfile)
	// Warning - log for fatal errors
	Warning = lg.New(os.Stderr, "[WARNING] ", lg.Lshortfile)
)
