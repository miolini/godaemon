package godaemon

import "testing"
import "os"

func TestWritePidFile(t *testing.T) {
    filename := "/tmp/godaemon.pid"
    WritePid(filename)
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        t.Errorf("pid not found: %s", filename)
        return
    }
    os.Remove(filename)
}

