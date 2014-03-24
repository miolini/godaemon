package godaemon

import "os"
import "strconv"
import "io/ioutil"

func WritePidFile(filename string) (pid int, err error) {
    pid = os.Getpid()
    err = ioutil.WriteFile(filename, []byte(strconv.Itoa(pid)), 0640)
    return
}

func RemovePidFile(filename string) (err error) {
    err = os.Remove(filename)
    return
}