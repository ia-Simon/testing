package main

import "regexp"

func main() {
	match, err := regexp.Match(`^record_.*\.log$`, []byte("record_2022-01-03T18:15:00.126234-03:00.log"))
	if err != nil {
		panic(err)
	}
	println(match)
}
