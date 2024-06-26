package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	pvKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	pubKey := &pvKey.PublicKey

	message := []byte("secret message")

	encMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, message, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", encMessage)

	signatureHashFunc := sha256.New()
	signatureHashFunc.Write(encMessage)
	msgHash := signatureHashFunc.Sum(nil)
	fmt.Printf("Encrypted hash: %x\n", msgHash)

	msgSignature, err := rsa.SignPSS(rand.Reader, pvKey, crypto.SHA256, msgHash, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted signature: %x\n", msgSignature)
}
