package main

import (
	"fmt"
)

func main() {
	test()

	fmt.Println("Hello World")
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	defer func() {
		panic("defer panic")
		fmt.Println("end of defer test")
	}()

	panic("normal panic")
	fmt.Println("end of test")
}
