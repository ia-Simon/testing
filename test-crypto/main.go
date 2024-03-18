package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
)

var apcClient *paymentcryptography.Client
var apcDataClient *paymentcryptographydata.Client

var operation = ""

func init() {
	flag.StringVar(&operation, "oper", operation, "Choose the operation to be performed by this test script")

	awsCfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("us-east-1"),
		config.WithSharedConfigProfile("asapcard-hml"),
	)
	if err != nil {
		panic(err)
	}

	apcClient = paymentcryptography.NewFromConfig(awsCfg)
	apcDataClient = paymentcryptographydata.NewFromConfig(awsCfg)
}

func main() {
	flag.Parse()
	ctx := context.Background()

	switch operation {
	case "DES":
		err := testDES()
		if err != nil {
			fmt.Printf("Failed DES operation. err=%v\n", err)
		}
	case "3DES_GetKCV":
		err := generate3DESKCV()
		if err != nil {
			fmt.Printf("Failed 3DES_KCV operation. err=%v\n", err)
		}
	case "3DES_3CompMerge":
		err := merge3DESKeyComponents(3)
		if err != nil {
			fmt.Printf("Failed 3DES_3CompMerge operation. err=%v\n", err)
		}
	case "3DES_2KeyHSMImport":
		err := import3DES2KeyToHSM(ctx)
		if err != nil {
			fmt.Printf("Failed 3DES_2KeyHSMImport operation. err=%v\n", err)
		}
	}
}

func testDES() error {
	key := make([]byte, des.BlockSize)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}

	blk, err := des.NewCipher(key)
	if err != nil {
		return err
	}

	fmt.Print("Message to be encrypted/decrypted with DES: ")
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	msg = strings.TrimRight(msg, "\n")

	msgBytes := []byte(msg)
	if len(msgBytes) < des.BlockSize {
		msgBytes = append(msgBytes, bytes.Repeat([]byte(" "), des.BlockSize-len(msgBytes))...)
	}

	encMsg := make([]byte, len(msgBytes))
	blk.Encrypt(encMsg, msgBytes)
	fmt.Printf("Encrypted: '%x'\n", encMsg)

	decMsg := make([]byte, len(msgBytes))
	blk.Decrypt(decMsg, encMsg)
	fmt.Printf("Decrypted: '%s'\n", decMsg)

	return nil
}

func generate3DESKCV() error {
	fmt.Print("3DES key (hex representation): ")
	reader := bufio.NewReader(os.Stdin)
	keyInput, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	keyInput = strings.TrimRight(keyInput, "\n")

	key, err := hex.DecodeString(keyInput)
	if err != nil {
		return err
	}

	switch len(key) {
	case 16:
		fmt.Println("Key with 16 bytes, mirroring K1 as K3...")
		key = append(key, key[:8]...)
	case 24:
		fmt.Println("Complete key with 24 bytes provided, using as is...")
	default:
		return fmt.Errorf("invalid key length %d", len(key))
	}

	blk, err := des.NewTripleDESCipher(key)
	if err != nil {
		return err
	}

	msg := bytes.Repeat([]byte{0x00}, des.BlockSize)
	encMsg := make([]byte, des.BlockSize)
	blk.Encrypt(encMsg, msg)
	fmt.Printf("%x\n", encMsg[:3])

	return nil
}

func merge3DESKeyComponents(componentCount int) error {
	reader := bufio.NewReader(os.Stdin)

	key := make([]byte, 24)
	for i := 0; i < componentCount; i++ {
		fmt.Printf("3DES key component %d (hex representation): ", i+1)
		keyCompInput, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		keyCompInput = strings.TrimRight(keyCompInput, "\n")

		keyComp, err := hex.DecodeString(keyCompInput)
		if err != nil {
			return err
		}

		switch len(keyComp) {
		case 16:
			fmt.Println("Key component with 16 bytes, mirroring K1 as K3...")
			keyComp = append(keyComp, keyComp[:8]...)
		case 24:
			fmt.Println("Complete key component with 24 bytes provided, using as is...")
		default:
			return fmt.Errorf("invalid component key length %d", len(keyComp))
		}

		for idx := range keyComp {
			key[idx] = key[idx] ^ keyComp[idx]
		}
	}

	fmt.Printf("2-key: %x\n", key[:16])
	fmt.Printf("3-key: %x\n", key)

	return nil
}

func import3DES2KeyToHSM(ctx context.Context) error {
	fmt.Print("3DES double length key (hex representation): ")
	reader := bufio.NewReader(os.Stdin)
	keyInput, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	keyInput = strings.TrimRight(keyInput, "\n")

	key, err := hex.DecodeString(keyInput)
	if err != nil {
		return err
	}
	if len(key) != 16 {
		return fmt.Errorf("invalid key length %d", len(key))
	}

	importParamsResp, err := apcClient.GetParametersForImport(ctx, &paymentcryptography.GetParametersForImportInput{
		KeyMaterialType:      types.KeyMaterialTypeKeyCryptogram,
		WrappingKeyAlgorithm: types.KeyAlgorithmRsa4096,
	})
	if err != nil {
		return err
	}

	certBytes, err := base64.StdEncoding.DecodeString(*importParamsResp.WrappingKeyCertificate)
	if err != nil {
		return err
	}
	certBlock, _ := pem.Decode(certBytes)
	if certBlock == nil {
		return errors.New("no PEM block found in APC generated wrapping certificate")
	}
	wrappingKeyCert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return err
	}

	wrappingPubKey, ok := wrappingKeyCert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("unable to extract RSA public key from APC generated wrapping certificate")
	}
	wrappedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, wrappingPubKey, key, nil)
	if err != nil {
		return err
	}
	wrappedKeyHex := strings.ToUpper(hex.EncodeToString(wrappedKey))

	importKeyResp, err := apcClient.ImportKey(ctx, &paymentcryptography.ImportKeyInput{
		KeyMaterial: &types.ImportKeyMaterialMemberKeyCryptogram{
			Value: types.ImportKeyCryptogram{
				Exportable:  aws.Bool(false),
				ImportToken: importParamsResp.ImportToken,
				KeyAttributes: &types.KeyAttributes{
					KeyAlgorithm: types.KeyAlgorithmTdes2key,
					KeyClass:     types.KeyClassSymmetricKey,
					KeyModesOfUse: &types.KeyModesOfUse{
						Decrypt:        true,
						DeriveKey:      false,
						Encrypt:        true,
						Generate:       false,
						NoRestrictions: false,
						Sign:           false,
						Unwrap:         true,
						Verify:         false,
						Wrap:           true,
					},
					KeyUsage: types.KeyUsageTr31K0KeyEncryptionKey,
				},
				WrappedKeyCryptogram: aws.String(wrappedKeyHex),
				WrappingSpec:         types.WrappingKeySpecRsaOaepSha256,
			},
		},
	})
	if err != nil {
		return err
	}

	createAliasResp, err := apcClient.CreateAlias(ctx, &paymentcryptography.CreateAliasInput{
		AliasName: aws.String(fmt.Sprintf("alias/test/KEK_%s", *importKeyResp.Key.KeyCheckValue)),
		KeyArn:    importKeyResp.Key.KeyArn,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Key Alias | ARN: %s | %s\n", *createAliasResp.Alias.AliasName, *importKeyResp.Key.KeyArn)
	fmt.Printf("Key KCV | Algo: %s | %s\n", *importKeyResp.Key.KeyCheckValue, importKeyResp.Key.KeyCheckValueAlgorithm)
	fmt.Printf("Key Type: %s\n", importKeyResp.Key.KeyAttributes.KeyAlgorithm)

	return nil
}
