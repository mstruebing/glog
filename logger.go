// Package logger provides ...
package glog

import (
	"fmt"
	"io"
)

type Logger struct {
	Target io.Writer
}

func (l Logger) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(l.Target, a)
}

func (l Logger) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(l.Target, format, a...)
}

func (l Logger) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(l.Target, a...)
}
