package main

import (
	"github.com/streadway/amqp"
	"log"
)

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

	err = ch.Publish(
		"AMQP", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("message"),
		},
	)

	log.Printf(" [x] Sent %s", "message")
}

func errorHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
