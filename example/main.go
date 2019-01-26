// Package main provides ...
package main

import (
	"github.com/mstruebing/glog"
	"github.com/mstruebing/glog/console"
	"github.com/mstruebing/glog/file"
	"os"
)

func main() {
	// console logging
	console.Error("This logs to stderr too")
	console.Log("This logs to stdout too")

	// file logging
	fileLogger := file.Logger{Target: "/tmp/mystuff.txt"}
	fileLogger.Println("This logs to a file")
	fileLogger.Printf("This logs to a %s\n", "file")

	file.Log("/tmp/mystuff.txt", "Another file log")

	// anywhere logging (io.Writer needed)
	genericLogger := glog.Logger{Target: os.Stdout}
	genericLogger.Println("Logs to stdout from genericLogger")
}
