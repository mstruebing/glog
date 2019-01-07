# WIP

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
