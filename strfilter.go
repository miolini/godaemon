package godaemon

func WorkerStringFilter(worker *Worker, inputChan chan string, outputChan chan string, fn func(string)(string,error)) (err error) {
    var str string
    for {
        str = <-inputChan
        str, err = fn(str)
        if err != nil {
            return
        }
        outputChan <- str
    }
    return
}
