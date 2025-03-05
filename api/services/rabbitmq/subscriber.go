package services

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqSubscriber struct{}

func (sub *RabbitMqSubscriber) Subscribe(queue string) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Print(string(d.Body))
		}
	}()

	<-forever
}
