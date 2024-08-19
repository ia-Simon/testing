package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
)

var (
	comp1Hex = ""
	comp2Hex = ""
	comp3Hex = ""

	// 2 for double length, 3 for triple length
	tdesLength = 3

	comps [][]byte
)

func init() {
	comp1, err := hex.DecodeString(comp1Hex)
	if err != nil {
		panic(err)
	}
	comps = append(comps, comp1)

	comp2, err := hex.DecodeString(comp2Hex)
	if err != nil {
		panic(err)
	}
	comps = append(comps, comp2)

	comp3, err := hex.DecodeString(comp3Hex)
	if err != nil {
		panic(err)
	}
	comps = append(comps, comp3)
}

func main() {
	key := make([]byte, des.BlockSize*tdesLength)
	for _, comp := range comps {
		for idx := range key {
			key[idx] ^= comp[idx]
		}
	}

	var tdesBlock cipher.Block
	var err error
	switch tdesLength {
	case 2:
		tdesBlock, err = des.NewTripleDESCipher(append(key, key[:8]...))
	case 3:
		tdesBlock, err = des.NewTripleDESCipher(key)
	default:
		err = errors.New("invalid tdesLength value")
	}
	if err != nil {
		panic(err)
	}

	kcv := make([]byte, des.BlockSize)
	tdesBlock.Encrypt(kcv, make([]byte, des.BlockSize))

	fmt.Printf("Key material: %X\nKCV: %X\n", key, kcv[:3])
}
