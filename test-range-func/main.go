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

	for item := range loop([]int{1, 2, 3, 4, 5}, 2) {
		fmt.Println("Item:", item)
		time.Sleep(150 * time.Millisecond)
	}
}

func loop[V any](slice []V, times int) func(func(V) bool) {
	return func(yield func(V) bool) {
		sliceLen := len(slice)
		loopedTimes := 0

		for ptr := sliceLen; ; ptr++ {
			if ptr == 2*sliceLen {
				ptr = sliceLen
				loopedTimes++
				if times > 0 && loopedTimes == times {
					return
				}
			}
			if !yield(slice[ptr%sliceLen]) {
				fmt.Println("stop iteration")
				break
			}
		}
	}
}
