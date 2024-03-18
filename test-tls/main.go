package main

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	caCertPath     string
	clientCertPath string
	clientKeyPath  string
	mipAddr        string
)

func init() {
	flag.StringVar(&caCertPath, "caCert", "", "CA certificate file path [.pem]")
	flag.StringVar(&clientCertPath, "clientCert", "", "Client certificate file path [.pem]")
	flag.StringVar(&clientKeyPath, "clientKey", "", "Client private key file path [.pem]")
	flag.StringVar(&mipAddr, "addr", "localhost:3125", "TLS server address")

	flag.Parse()

	if caCertPath == "" {
		log.Fatal("[flag] caCert path must be specified")
	}
	if clientCertPath == "" {
		log.Fatal("[flag] clientCert path must be specified")
	}
	if clientKeyPath == "" {
		log.Fatal("[flag] clientKey path must be specified")
	}
}

func main() {
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Failed to read CA certificate. err='%v'", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		log.Fatalf("Failed to load client certificate/key pair. err='%v'", err)
	}

	conn, err := tls.Dial("tcp", mipAddr, &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatalf("Failed to dial TLS connection. err='%v'", err)
	}

	log.Print("Listening to connection...")
	readBuffer := make([]byte, 1024)
	for {
		lenRead, err := conn.Read(readBuffer)
		if err != nil {
			log.Fatalf("Failed to read from TLS connection. err='%v'", err)
		}
		fmt.Printf("Read buffer: %x", readBuffer[:lenRead])

		if lenRead == 2 && binary.BigEndian.Uint16(readBuffer[:lenRead]) == 0 {
			zeroLengthResponse := []byte{0x00, 0x00}

			_, err := conn.Write(zeroLengthResponse)
			if err != nil {
				log.Fatalf("Failed to write to TLS connection. err='%v'", err)
			}

			continue
		}

		// Insert message handling code here
	}
}
