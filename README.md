# Glog - a golang logging library made for simplicity
[![Build Status](https://travis-ci.org/mstruebing/glog.svg?branch=master)](https://travis-ci.org/mstruebing/glog)
[![codecov](https://codecov.io/gh/mstruebing/glog/branch/master/graph/badge.svg)](https://codecov.io/gh/mstruebing/glog)
[![Go Report Card](https://goreportcard.com/badge/github.com/mstruebing/glog)](https://goreportcard.com/report/github.com/mstruebing/glog)

## This is currently WIP

Look at `./example` for what it is doing.


```go
// Package main provides ...
package main

import (
	"github.com/mstruebing/glog"
	"github.com/mstruebing/glog/console"
	"github.com/mstruebing/glog/file"
	"os"
)

func main() {
	stderrLogger := console.StderrLogger()
	stderrLogger.Println("This logs to stderr")

	stdoutLogger := console.StdoutLogger()
	stdoutLogger.Println("This logs to stdout")

	fileLogger := file.Logger{Target: "/tmp/mystuff.txt"}
	fileLogger.Println("This logs to a file")

	genericLogger := glog.Logger{Target: os.Stdout}
	genericLogger.Println("Logs to stdout from genericLogger")
}
```
