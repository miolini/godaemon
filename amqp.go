package godaemon

import "log"
import "time"
import "github.com/streadway/amqp"

func WorkerAmqpConsumer(worker *Worker, uri string, queue string, msgChan chan string) (err error) {
    queueConn, err := amqp.Dial(uri)
    if err != nil {
        return
    }
    defer queueConn.Close()
    queueChan, err := queueConn.Channel()
    if err != nil {
        return
    }
    defer queueChan.Close()
    msgs, err := queueChan.Consume(queue, "", true, false, false, false, nil)
    for {
        for len(msgChan) > cap(msgChan) / 100 * 95 {
            log.Printf("worker amqp consumer wait busy channel: %d/%d", len(msgChan), cap(msgChan))
            time.Sleep(time.Millisecond * 100)
        }
        msg := <- msgs
        msgChan <- string(msg.Body)
    }
    return
}

func WorkerAmqpPublisher(worker *Worker, uri string, exchange string, msgChan chan string) (err error) {
    queueConn, err := amqp.Dial(uri)
    if err != nil {
        return
    }
    defer queueConn.Close()
    queueChan, err := queueConn.Channel()
    if err != nil {
        return
    }
    defer queueChan.Close()
    for msg := range msgChan {
        publishing := amqp.Publishing{ContentType:"text/plain",Body:[]byte(msg)}
        queueChan.Publish(exchange, "", false, false, publishing)
    }
    return
}