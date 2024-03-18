package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.DialTLS("amqp://guest:guest@0.tcp.ngrok.io:17644/", nil)
	if err != nil {
		log.Fatalln("failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel", err)
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

	if err = ch.Qos(1, 0, false); err != nil {
		log.Fatalln("failed to set QoS", err)
	}

	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to register as consumer", err)
	}

	for msg := range msgs {
		t_processDelivery(msg)
	}
}

func t_processDelivery(d amqp.Delivery) (err error) {
	defer d.Ack(false)

	log.Printf("JSON log received: %s", string(d.Body))

	return
}
