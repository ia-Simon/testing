package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		panic(fmt.Sprintf("test returned error: %v", err))
	}
}

func test() (err error) {
	defer func() {
		if err != nil {
			fmt.Println("defer found error")
			err = nil
		}
	}()

	err = mustError()
	if err != nil {
		return err
	}

	return nil
}

func mustError() error {
	return errors.New("must error")
}
