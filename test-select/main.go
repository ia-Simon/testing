package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		a := <-ch
		fmt.Println(a)
	}()

loop1:
	for {
		select {
		case ch <- "1":
			fmt.Println("Sent 1")
			break loop1
		default:
			fmt.Println("Not sent 1")
			time.Sleep(10 * time.Millisecond)
		}
	}

loop2:
	for {
		select {
		case ch <- "2":
			fmt.Println("Sent 2")
			break loop2
		default:
			fmt.Println("Not sent 2")
			time.Sleep(10 * time.Millisecond)
		}
	}
}
