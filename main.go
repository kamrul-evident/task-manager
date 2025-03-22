package main

import (
    "log"
    "time"
    "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

// Producer function
func publishMessage() {
    // Connect to RabbitMQ server
    conn, err := amqp091.Dial("amqp://kamrul:kamrul@localhost:5674/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // Create a channel
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // Declare a queue
    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    // Message to send
    body := "Hello World!"
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp091.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    failOnError(err, "Failed to publish a message")
    log.Printf(" [x] Sent %s", body)
}

// Consumer function
func consumeMessages() {
    conn, err := amqp091.Dial("amqp://kamrul:kamrul@localhost:5674/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    // Consume messages
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever
}

func main() {
    // Run producer in a goroutine
    go func() {
        for {
            publishMessage()
            time.Sleep(2 * time.Second) // Send every 2 seconds
        }
    }()

    // Run consumer
    consumeMessages()
}