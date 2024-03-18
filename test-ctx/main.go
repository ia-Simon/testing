package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	go func() {
		t := 5 * time.Second
		fmt.Printf("Sleeping %v...\n", t)
		time.Sleep(t)
		fmt.Println("Done sleeping")
		cancelFunc()
	}()

	<-ctx.Done()
}
