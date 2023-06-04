package main

import (
	"log"
	"sync"

	"github.com/samurainet-antonius/rabbit-libs/lib/debug"
	"github.com/samurainet-antonius/rabbit-libs/rabbitmq"
)

func main() {
	debug.Debug = true

	conn, err := rabbitmq.Dial("amqp://rabbit:passrabbit@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}

	exchangeName := "test-exchange"
	queueName := "test-queue"
	key := "test.log*"

	consumeCh, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	consumeCh.Qos(1, 0, false)
	q, err := consumeCh.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}

	if err := consumeCh.QueueBind(q.Name, key, exchangeName, false, nil); err != nil {
		log.Panic(err)
	}

	go func() {
		d, err := consumeCh.Consume(q.Name, "", false, false, false, false, nil)
		if err != nil {
			log.Panic(err)
		}

		for msg := range d {
			log.Printf("msg: %s", string(msg.Body))
			msg.Ack(true)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
}
