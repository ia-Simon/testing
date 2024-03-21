package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	useTLS   bool
	addr     = "localhost:6379"
	password string
	db       int
)

func init() {
	flag.BoolVar(&useTLS, "tls", useTLS, "if set, connection with Redis will be made using TLS")
	flag.StringVar(&addr, "addr", addr, "redis address")
	flag.StringVar(&password, "pass", password, "redis password")
}

func main() {
	flag.Parse()

	var conn *redis.Client
	if useTLS {
		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}

		conn = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
			TLSConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		})
	} else {
		conn = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
	}

	res, err := conn.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
