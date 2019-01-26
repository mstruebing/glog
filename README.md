# Glog - a golang logging library made for simplicity
[![Build Status](https://travis-ci.org/mstruebing/glog.svg?branch=master)](https://travis-ci.org/mstruebing/glog)
[![codecov](https://codecov.io/gh/mstruebing/glog/branch/master/graph/badge.svg)](https://codecov.io/gh/mstruebing/glog)
[![Go Report Card](https://goreportcard.com/badge/github.com/mstruebing/glog)](https://goreportcard.com/report/github.com/mstruebing/glog)

## This is currently WIP

Look at `./example` for what it is doing.

```
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
```
