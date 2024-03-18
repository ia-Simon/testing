package main

import (
	"crypto/tls"
	"log"

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

	queue, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatalln("failed to declare a queue")
	}

	err = ch.QueueBind(queue.Name, "", "logs", false, nil)
	if err != nil {
		log.Fatalln("failed to bind a queue")
	}

	if err = ch.Qos(100, 0, false); err != nil {
		log.Fatalln("failed to set QoS")
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to register as consumer")
	}

	for msg := range msgs {
		s_processDelivery(msg)
	}
}

func s_processDelivery(d amqp.Delivery) (err error) {
	switch d.ContentType {
	case "application/json":
		log.Printf("JSON log received: %s", string(d.Body))
	}

	return
}
