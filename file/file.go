// Package file provides ...
package file

import (
	"bufio"
	"github.com/mstruebing/glog"
	"os"
)

// Logger is different than the generice one.
// Here we are using a string to indicate the filePath
// instead of an io.Writer
type Logger struct {
	Target string
}

// Print appends a string to a file in the form of fmt.Print
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

// Printf appends a string to a file in the form of fmt.Printf
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

// Println appends a line to a file in the form of fmt.Println
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

// Log appends the given string a to the filePath
func Log(filePath string, a ...interface{}) (int, error) {
	f, err := openOrCreateFile(filePath)
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
