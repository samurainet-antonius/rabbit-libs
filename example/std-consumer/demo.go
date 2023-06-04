package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	consumer "git.ainosi.co.id/go-libs/rabbit-lib/consumer"
	"git.ainosi.co.id/go-libs/rabbit-lib/lib/debug"
	"git.ainosi.co.id/go-libs/rabbit-lib/rabbitmq"
	config "github.com/joho/godotenv"
	"github.com/streadway/amqp"
	// publisher "git.ainosi.co.id/go-libs/rabbit-lib/publisher"
)

func main() {
	debug.Debug = true

	if err := config.Load("env/.env"); err != nil {
		fmt.Println(".env is not loaded properly")
		fmt.Println(err)
		os.Exit(2)
	}

	conn, err := rabbitmq.Dial(os.Getenv("RABBIT_URL"))
	if err != nil {
		log.Panic(err)
	}

	exchangeName := "test-exchange-topic"
	// exchangeName := "consolidation.transaction"
	queueName := "test-queue-topic"
	key := []string{"test.log.coba2"}

	consumeCh, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	consumeCh.Qos(1, 0, false)

	cons := consumer.NewConsumer(consumeCh, exchangeName, key, queueName)

	// add handler to consume message
	cons.ConsumeMessage(HandlerMessage)

	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
}

// HandlerMessage :
func HandlerMessage(msg amqp.Delivery) {
	log.Printf("msg: %s %s , header: %s", string(msg.Body), "read", msg.Headers["a"])
}
