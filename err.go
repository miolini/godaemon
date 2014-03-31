package godaemon

import "log"

func CheckErr (label string, err error) {
    if err != nil {
        log.Fatalf("error %s: %s", label, err)
    }
}