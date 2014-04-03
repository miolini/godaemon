package godaemon

import "sync"
import "time"
import "log"

func NewStatCounter(label string, interval int) (chan int) {
    statChan := make(chan int)
    counter := 0
    timer := time.Now()
    mutex := sync.Mutex{}
    go func() {
        for volume := range statChan {
            mutex.Lock()
            counter += volume
            mutex.Unlock()
        }
    }()
    go func () {
        for {
            mutex.Lock()
            log.Printf("speed %s %d msg/sec", label, counter / interval)
            counter = 0
            timer = time.Now()
            mutex.Unlock()
            time.Sleep(time.Second * time.Duration(interval))
        }
    }()
    return statChan
}