package main

import (
	"bytes"
	"context"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
)

var (
	pvkAlias = "alias/test-PVK"
	pekAlias = "alias/test-PEK"

	pekComps = []string{
		"",
		"",
		"",
	}

	pin = "1234"
	pan = "1234567800000001"
)

func main() {
	ctx := context.Background()

	awsConf, err := config.LoadDefaultConfig(ctx,
		config.WithDefaultRegion("us-east-1"),
		config.WithSharedConfigProfile("default"),
	)
	if err != nil {
		panic(err)
	}

	apcDataClient := paymentcryptographydata.NewFromConfig(awsConf)

	pek := make([]byte, 24)
	for _, comp := range pekComps {
		binComp, err := hex.DecodeString(comp)
		if err != nil {
			panic(err)
		}

		for i := range pek {
			pek[i] ^= binComp[i]
		}
	}

	pekTDESBlock, err := des.NewTripleDESCipher(pek)
	if err != nil {
		panic(err)
	}
	kcv := make([]byte, des.BlockSize)
	pekTDESBlock.Encrypt(kcv, bytes.Repeat([]byte{0x00}, des.BlockSize))
	fmt.Printf("PEK: %s\n", strings.ToUpper(hex.EncodeToString(pek)))
	fmt.Printf("KCV: %s\n", strings.ToUpper(hex.EncodeToString(kcv[:3])))

	// ISO format 0 PIN
	// See https://www.eftlab.com/knowledge-base/complete-list-of-pin-blocks#AS2805-1 for more information

	if len(pin) < 4 || len(pin) > 12 {
		panic(fmt.Sprintf("invalid pin length %d", len(pin)))
	}
	paddedPin := fmt.Sprintf("0%x%s", len(pin), pin)
	paddedPin += strings.Repeat("F", 16-len(paddedPin))
	fmt.Println("Padded PIN:", paddedPin)

	if len(pan) < 12 || len(pan) > 19 {
		panic(fmt.Sprintf("invalid pan length %d", len(pan)))
	}
	paddedPan := fmt.Sprintf("0000%s", pan[len(pan)-13:len(pan)-1])
	fmt.Println("Padded PAN:", paddedPan)

	bytesPaddedPin, err := hex.DecodeString(paddedPin)
	if err != nil {
		panic(err)
	}

	bytesPaddedPan, err := hex.DecodeString(paddedPan)
	if err != nil {
		panic(err)
	}

	pinBlock := make([]byte, 8)
	for idx := range pinBlock {
		pinBlock[idx] = bytesPaddedPin[idx] ^ bytesPaddedPan[idx]
	}

	fmt.Println("PIN Block:", strings.ToUpper(hex.EncodeToString(pinBlock)))

	encryptedPinBlock := make([]byte, 8)
	pekTDESBlock.Encrypt(encryptedPinBlock, pinBlock)

	fmt.Println("Encrypted PIN Block:", strings.ToUpper(hex.EncodeToString(encryptedPinBlock)))

	genPinResp, err := apcDataClient.GeneratePinData(ctx, &paymentcryptographydata.GeneratePinDataInput{
		EncryptionKeyIdentifier: aws.String(pekAlias),
		GenerationKeyIdentifier: aws.String(pvkAlias),
		PinBlockFormat:          types.PinBlockFormatForPinDataIsoFormat0,
		PrimaryAccountNumber:    aws.String(pan),
		GenerationAttributes: &types.PinGenerationAttributesMemberVisaPinVerificationValue{
			Value: types.VisaPinVerificationValue{
				EncryptedPinBlock:       aws.String(hex.EncodeToString(encryptedPinBlock)),
				PinVerificationKeyIndex: aws.Int32(1),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	pinData, ok := genPinResp.PinData.(*types.PinDataMemberVerificationValue)
	if !ok {
		panic("unexpected pin generation response")
	}

	fmt.Printf("Pin Verification Value: %s\n", pinData.Value)
}
