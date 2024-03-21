package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"net/url"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	useTLS   bool
	host     = "localhost"
	port     = "5672"
	user     = "guest"
	password = "guest"
	db       int
)

func init() {
	flag.BoolVar(&useTLS, "tls", useTLS, "if set, connection with RabbitMQ will be made using TLS")
	flag.StringVar(&host, "host", host, "rabbitMQ host")
	flag.StringVar(&port, "port", port, "rabbitMQ port")
	flag.StringVar(&user, "user", user, "rabbitMQ user")
	flag.StringVar(&password, "pass", password, "rabbitMQ password")
}

func main() {
	flag.Parse()

	schema := "amqp"
	if useTLS {
		schema = "amqps"
	}

	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}

	conn, err := amqp.DialTLS(
		fmt.Sprintf("%s://%s:%s@%s:%s/", schema, user, url.QueryEscape(password), host, port),
		&tls.Config{
			RootCAs: caCertPool,
		},
	)
	if err != nil {
		panic(err)
	}

	_, err = conn.Channel()
	if err != nil {
		panic(err)
	}
}
