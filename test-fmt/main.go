package main

import "fmt"

func main() {
	padding := 2

	fmt.Printf(fmt.Sprintf(">>%%0%dd\n", padding), 1)

	fmt.Printf("%02s %02d", "123", 3)
}
