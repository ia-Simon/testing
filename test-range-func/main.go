package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"aasd", "a23", "hgh45", "h6ds42", "gfd943h"}

	for idx, d := range data {
		fmt.Println("Index:", idx, "\n", "Data:", d)
	}

	fmt.Println()

	for num := range 5 {
		fmt.Println("Count:", num)
	}

	fmt.Println()

	ch := time.After(10 * time.Second)
loop:
	for item := range loop([]int{1, 2, 3, 4, 5}) {
		fmt.Println("Item:", item)
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ch:
			break loop
		default:
		}
	}
}

func loop[V any](slice []V) func(func(V) bool) {
	return func(yield func(V) bool) {
		sliceLen := len(slice)
		for ptr := sliceLen; ; ptr++ {
			if ptr == 2*sliceLen {
				ptr = sliceLen
			}
			if !yield(slice[ptr%sliceLen]) {
				fmt.Println("stop iteration")
				break
			}
		}
	}
}
