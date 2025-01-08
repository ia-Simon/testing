package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

var file = flag.String("file", "", "File to count words in")

func main() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	flag.Parse()

	if *file == "" {
		panic("Please provide a file to count words in using the -file flag")
	}

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	buff := bufio.NewReaderSize(f, 4096)

	count := 0
	inWord := false
	for contentAvailable := true; contentAvailable; {
		b, err := readByte(buff)
		if err != nil {
			if err == io.EOF {
				contentAvailable = false
			} else {
				panic(err)
			}
		}

		if unicode.IsSpace(rune(b)) && inWord {
			count++
		}

		inWord = !unicode.IsSpace(rune(b))
	}

	fmt.Printf("Word count: %d\n", count)
}

var buf [1]byte

func readByte(r io.Reader) (byte, error) {
	_, err := r.Read(buf[:])
	if err != nil {
		return 0, err
	}

	return buf[0], nil
}
