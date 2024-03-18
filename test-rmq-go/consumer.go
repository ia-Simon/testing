package main

import (
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

	queue, err := ch.QueueDeclare("hello-world-queue", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to declare a queue")
	}

	if err = ch.Qos(1, 0, false); err != nil {
		log.Fatalln("failed to set QoS")
	}

	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to register as consumer")
	}

	for msg := range msgs {
		c_processDelivery(msg)
	}
}

func c_processDelivery(d amqp.Delivery) (err error) {
	defer d.Ack(false)

	log.Printf("processing message #%s", d.MessageId)

	switch d.ContentType {
	case "application/json":
		log.Printf("JSON message received: %s", string(d.Body))
		data := map[string]any{}
		err = json.Unmarshal(d.Body, &data)

		time.Sleep(time.Second * time.Duration(data["complexity"].(float64)))

		log.Printf("JSON message processed: %+v", data)
	}

	return
}
