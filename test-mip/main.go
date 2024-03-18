package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	READ_BUFFER_SIZE  = 1024
	READ_RESP_TIMEOUT = 3
)

func main() {
	l, err := net.Listen("tcp", ":3125")
	if err != nil {
		panic(err)
	}

	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}

	readBuffer := make([]byte, READ_BUFFER_SIZE)
	for {
		msg := readMessage()

		conn.Write(msg)

		conn.SetReadDeadline(time.Now().Add(time.Second * READ_RESP_TIMEOUT))
		lenRead, err := conn.Read(readBuffer)
		if err != nil {
			slog.Error("Failed to read response", slog.Any("err", err))
		}

		fmt.Println(string(readBuffer[:lenRead]))
	}
}

const Sep = `\x`

func readMessage() []byte {
	fmt.Printf("Insira e envie com <enter>:\n")

	// Lê o input (mensagem)
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.ReplaceAll(msg, "\n", "")

	// Insere separador
	var message string
	var sepCounter int64
	for i := 0; i < len(msg); i += 2 {
		sepCounter++
		message += string(Sep) + string(msg[i])
		// Prevê panic por index out of range
		if i+1 >= len(msg) {
			break
		}
		message += string(msg[i+1])
	}

	sepCounterToHex := strconv.FormatInt(sepCounter, 16)
	msgPrefix := fmt.Sprintf("%04s", sepCounterToHex)

	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("Resultado:")

	// Escreve a mensagem final em STDOUT
	isoMessage := Sep + msgPrefix[0:2] + Sep + msgPrefix[2:4] + message
	fmt.Println(isoMessage)

	isoBytes, err := hex.DecodeString(strings.ReplaceAll(isoMessage, Sep, ""))
	if err != nil {
		slog.Error("Error decoding message from string", slog.Any("err", err))
	}
	return isoBytes
}
