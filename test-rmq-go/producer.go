package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.DialTLS("amqp://guest:guest@localhost:5672/", nil)
	if err != nil {
		log.Fatalln("failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel")
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("hello-world-queue", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to declare a queue")
	}

	payload, err := json.Marshal(map[string]any{
		"message":    "hello world!",
		"complexity": rand.Intn(5),
	})
	if err != nil {
		log.Fatalln("failed to marshal payload")
	}

	err = ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         payload,
		MessageId:    uuid.New().String(),
	})

	if err != nil {
		log.Fatalln("failed to publish a message")
	}

	log.Printf("message '%s' sent successfully", payload)
}
