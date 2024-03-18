package main

import (
	"fmt"
	"os"
)

func main() {
	resp, err := os.ReadDir("/")
	if err != nil {
		panic(err)
	}

	for _, v := range resp {
		fmt.Printf("%+v\n", v)
	}
}
