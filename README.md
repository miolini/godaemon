godaemon
========

Simple Daemonize Helper Code


Exmaple

```
package main

import "github.com/miolini/godaemon"
import "time"
import "log"

func main() {
    pidFile := "/var/run/mydaemon.pid"
    err := godaemon.WritePidFile(pidFile)
    if err != nil {
        log.Fatalf("pid error: %s", err)
    }
    time.Sleep(5)
    godaemon.RemovePidFile(pidFile)
}
```


