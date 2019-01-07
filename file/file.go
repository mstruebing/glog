// Package file provides ...
package file

import (
	"bufio"
	"github.com/mstruebing/glog"
	"os"
)

type Logger struct {
	Target string
}

func (l Logger) Print(a ...interface{}) (int, error) {
	f, err := openOrCreateFile(l.Target)
	defer f.Close()

	if err != nil {
		return 0, err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	logger := glog.Logger{Target: w}

	return logger.Print(a...)
}

func (l Logger) Printf(format string, a ...interface{}) (int, error) {
	f, err := openOrCreateFile(l.Target)
	defer f.Close()

	if err != nil {
		return 0, err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	logger := glog.Logger{Target: w}

	return logger.Printf(format, a...)
}

func (l Logger) Println(a ...interface{}) (int, error) {
	f, err := openOrCreateFile(l.Target)
	defer f.Close()

	if err != nil {
		return 0, err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	logger := glog.Logger{Target: w}

	return logger.Println(a...)
}

func openOrCreateFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		f, err := os.Create(path)
		return f, err
	}

	return f, err
}
