package services

import (
	"log"
	"time"
	"github.com/rabbitmq/amqp091-go"
)

// func failOnError (err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

func publishMessage() {
	// Connect to rabbit mq server
	conn, err := amqp091.Dial("amqp://kamrul:kamrul@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil, //arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Message to send
	body := "Hello World!"
	err = ch.Publish(
		"", //exchange
		q.Name, // routing key
		false, //mandatory
		false, //immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body: []byte(body)
		}
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
