// Package logger provides ...
package glog

import (
	"fmt"
	"io"
)

// Logger struct which can have a Target
type Logger struct {
	Target io.Writer
}

// Print prints to the loggers Target like fmt.Print
func (l Logger) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(l.Target, a...)
}

// Printf prints to the loggers Target like fmt.Printf does
func (l Logger) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(l.Target, format, a...)
}

// Println prints to the loggers Target like fmt.Println does
func (l Logger) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(l.Target, a...)
}
