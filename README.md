godaemon
========

Simple Daemonize Helper Code


Exmaple

```
import "github.com/miolini/godaemon"

func main() {
    pidFile := "/var/run/mydaemon.pid"
    godaemon.WritePidFile(pidFile)
    time.Sleep(5)
    godaemon.RemovePidFile(pidFile)
}
```


