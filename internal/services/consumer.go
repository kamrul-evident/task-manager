package services

import (
	"log"
	"time"
	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}