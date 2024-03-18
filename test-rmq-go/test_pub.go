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
	conn, err := amqp.DialTLS("amqp://guest:guest@0.tcp.ngrok.io:17644/", nil)
	if err != nil {
		log.Fatalln("failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("debezium.public.debezium_table", amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to declare exchange", err)
	}

	queue, err := ch.QueueDeclare("debezium_table_events", true, false, false, false, amqp.Table{"x-queue-type": "stream"})
	if err != nil {
		log.Fatalln("failed to declare queue", err)
	}

	err = ch.QueueBind(queue.Name, "debezium-key", "debezium.public.debezium_table", false, nil)
	if err != nil {
		log.Fatalln("failed to bind queue", err)
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
