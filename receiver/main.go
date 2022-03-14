package main

import (
	"github.com/streadway/amqp"
	"log"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errorHandler(err)
	defer conn.Close()

	ch, err := conn.Channel()
	errorHandler(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"AMQP",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	errorHandler(err)

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	errorHandler(err)

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"AMQP", // exchange
		false,
		nil,
	)
	errorHandler(err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	errorHandler(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
