package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()
	twoWeeksAfter := time.Date(now.Year(), now.Month(), now.Day()+14, 0, 0, 0, 0, now.Location())

	fmt.Println(now, twoWeeksAfter)
}
