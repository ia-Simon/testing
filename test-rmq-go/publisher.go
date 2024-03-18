package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.DialTLS("amqp://guest:guest@localhost:5672/", &tls.Config{})
	if err != nil {
		log.Fatalln("failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", amqp.ExchangeFanout, true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to declare an exchange")
	}

	payload, err := json.Marshal(map[string]any{
		"message":   "hello world!",
		"context":   map[string]any{},
		"timestamp": time.Now().Format(time.RFC3339),
	})
	if err != nil {
		log.Fatalln("failed to marshal payload")
	}

	err = ch.PublishWithContext(context.Background(), "logs", "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        payload,
	})

	if err != nil {
		log.Fatalln("failed to publish a message")
	}

	log.Printf("log '%s' sent successfully", payload)
}
