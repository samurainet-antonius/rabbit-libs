package consumer

import (
	"log"

	"github.com/samurainet-antonius/rabbit-libs/rabbitmq"
	"github.com/streadway/amqp"
)

type Consumer interface {
	ConsumeMessage(func(msg amqp.Delivery))
}

type consumer struct {
	ch           *rabbitmq.Channel
	exchangeName string
	routingKey   []string
	queueName    string
}

// NewConsumer :
func NewConsumer(ch *rabbitmq.Channel, exchangeName string, routingKeys []string, queueName string) Consumer {
	err := ch.ExchangeDeclare(exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}
	for _, routingKey := range routingKeys {
		if err := ch.QueueBind(q.Name, routingKey, exchangeName, false, nil); err != nil {
			log.Panic(err)
		}
	}

	return &consumer{
		ch:           ch,
		exchangeName: exchangeName,
		routingKey:   routingKeys,
		queueName:    queueName,
	}
}

func (cons *consumer) ConsumeMessage(process func(msg amqp.Delivery)) {
	go func() {
		d, err := cons.ch.Consume(cons.queueName, "", false, false, false, false, nil)
		if err != nil {
			log.Panic(err)
		}

		for msg := range d {
			// log.Printf("msg: %s", string(msg.Body))
			process(msg)
			msg.Ack(true)
		}
	}()
}
