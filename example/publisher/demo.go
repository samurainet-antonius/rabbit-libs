package main

import (
	"log"
	"sync"
	"time"

	"git.ainosi.co.id/go-libs/rabbit-lib/lib/debug"
	"git.ainosi.co.id/go-libs/rabbit-lib/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	debug.Debug = true

	conn, err := rabbitmq.Dial("amqp://rabbit:passrabbit@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}

	sendCh, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	exchangeName := "test-exchange"
	key := "test.log*"

	err = sendCh.ExchangeDeclare(exchangeName, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}

	// _, err = sendCh.QueueDeclare(queueName, true, false, false, false, nil)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// go func() {
	for {
		err := sendCh.Publish(exchangeName, key, true, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(time.Now().String()),
		})
		log.Printf("publish, err: %v", err)
		time.Sleep(2 * time.Second)
	}
	// }()

	// consumeCh, err := conn.Channel()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// go func() {
	// 	d, err := consumeCh.Consume(queueName, "", false, false, false, false, nil)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}

	// 	for msg := range d {
	// 		log.Printf("msg: %s", string(msg.Body))
	// 		msg.Ack(true)
	// 	}
	// }()

	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
}
