package godaemon

import "time"
import "sync"
import "log"

type Worker struct {
    Label string
    Num int
    Sleep time.Duration
    Func func(*Worker) error
    waitGroup sync.WaitGroup
    needStop bool
}

func NewWorker (label string, num int, sleep time.Duration) (worker *Worker) {
    worker = new(Worker)
    worker.Label = label
    worker.Num = num
    worker.Sleep = sleep
    return
}

func (worker *Worker) Start(fn func(*Worker) error) *Worker {
    worker.waitGroup = sync.WaitGroup{}
    worker.Func = fn
    worker.waitGroup.Add(worker.Num)
    for i:=0;i<worker.Num;i++ {
        go func() {
            var err error
            for {
                err = worker.Func(worker)
                if err != nil {
                    log.Printf("worker err (%s): %s", worker.Label, err)
                    if !worker.needStop {
                        time.Sleep(worker.Sleep)
                    }
                }
                if worker.needStop {
                    break
                }
            }
            worker.waitGroup.Done()
        }()
    }
    return worker
}

func (worker *Worker) Wait() {
    worker.waitGroup.Wait()
}

func (worker *Worker) Stop() {
    worker.needStop = true
    worker.Wait()
}