// Package console provides ...
package console

import (
	"fmt"
	"github.com/mstruebing/glog"
	"os"
)

const (
	YELLOW = "\x1b[33;1m"
	GREEN  = "\x1b[32;1m"
	RED    = "\x1b[31;1m"
	RESET  = "\x1b[33;0m"
)

// StdoutLogger creates a Logger which has the Target os.Stdout
func StdoutLogger() glog.Logger {
	return glog.Logger{Target: os.Stdout}
}

// StderrLogger creates a Logger which has the Target os.Stderr
func StderrLogger() glog.Logger {
	return glog.Logger{Target: os.Stderr}
}

// Log logs to os.Stdout
func Log(a ...interface{}) (int, error) {
	return fmt.Fprintln(os.Stdout, a...)
}

// Error logs to os.Stderr
func Error(a ...interface{}) (int, error) {
	return fmt.Fprintln(os.Stderr, a...)
}
