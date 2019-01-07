// Package console provides ...
package console

import (
	"github.com/mstruebing/glog"
	"os"
)

const (
	YELLOW = "\x1b[33;1m"
	GREEN  = "\x1b[32;1m"
	RED    = "\x1b[31;1m"
	RESET  = "\x1b[33;0m"
)

func StdoutLogger() glog.Logger {
	return glog.Logger{Target: os.Stdout}
}

func StderrLogger() glog.Logger {
	return glog.Logger{Target: os.Stderr}
}
