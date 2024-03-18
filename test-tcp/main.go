package main

import (
	"net"
	"sync"
	"time"
)

var msg = make([]byte, 0, 1*1024*1024)

func main() {
	msg = append(msg, []byte{
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
		0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56, 0x56,
	}...) // 64 bytes
	msg = append(msg, msg...) // x2
	msg = append(msg, msg...) // x4
	msg = append(msg, msg...) // x8
	msg = append(msg, msg...) // x16
	msg = append(msg, msg...) // x32
	msg = append(msg, msg...) // x64
	// msg = append(msg, msg...) // x128
	// msg = append(msg, msg...) // x256
	// msg = append(msg, msg...) // x512
	// msg = append(msg, msg...) // x1024
	// msg = append(msg, msg...) // x2048
	// msg = append(msg, msg...) // x4096
	// msg = append(msg, msg...) // x8192
	// msg = append(msg, msg...) // x16384
	msg = append(msg, 0x0a, 0x0a)

	msgs := make([][]byte, 0, 3)
	msgs = append(msgs, msg)
	msgs = append(msgs, msg)
	msgs = append(msgs, msg)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// connWriteLock := sync.Mutex{}
	targetTime := time.Now().Add(3 * time.Second)
	wg := &sync.WaitGroup{}

	for _, msg := range msgs {
		wg.Add(1)
		go func(data []byte) {
			<-time.After(time.Until(targetTime))

			// connWriteLock.Lock()
			_, err = conn.Write(data)
			// connWriteLock.Unlock()

			wg.Done()
		}(msg)
	}

	wg.Wait()
}
