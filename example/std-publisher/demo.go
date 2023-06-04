package main

import (
	"log"
	"time"

	rabbitPub "github.com/samurainet-antonius/rabbit-libs/publisher"
	rabbit "github.com/samurainet-antonius/rabbit-libs/rabbitmq"
)

func main() {
	var err error
	var conn *rabbit.Connection
	var ch *rabbit.Channel

	exchangeName := "test-exchange-topic"
	key := "test.log.coba2"

	conn, err = rabbit.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}
	ch, err = conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	pub := rabbitPub.NewStandartPublisher(ch, exchangeName)

	headers := make(map[string]interface{})
	headers["Content-type"] = "application/json"

	for i := 0; i <= 50; i++ {
		go publish(pub, key, headers)
	}

	forever := make(chan int)
	<-forever
}

func publish(pub rabbitPub.Publisher, key string, headers map[string]interface{}) {
	for {
		if err := pub.PublishMessage(key, []byte(`{"test":"just a test"}`), headers); err != nil {
			log.Println(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
