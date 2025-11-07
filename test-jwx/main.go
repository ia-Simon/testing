package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwe"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func main() {
	jws_key, err := os.ReadFile("jws_key")
	if err != nil {
		panic(err)
	}

	jwe_key, err := os.ReadFile("jwe_key")
	if err != nil {
		panic(err)
	}

	//############

	fmt.Println("JWT:")

	jwtToken := jwt.New()
	err = jwtToken.Set("exp", time.Now().UTC().Add(3600*time.Second))
	if err != nil {
		panic(err)
	}
	err = jwtToken.Set("custom1", "1234")
	if err != nil {
		panic(err)
	}

	jwtSigned, err := jwt.Sign(jwtToken, jwt.WithKey(jwa.HS512(), jws_key))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n%s\n", jwtToken, jwtSigned)

	//############

	fmt.Println("JWS:")

	jwsPayload := map[string]any{
		"exp":     time.Now().UTC().Add(3600 * time.Second).Unix(),
		"custom1": "5678",
	}
	jwsPayloadBytes, err := json.Marshal(jwsPayload)
	if err != nil {
		panic(err)
	}

	jwsSigned, err := jws.Sign(jwsPayloadBytes, jws.WithKey(jwa.HS512(), jws_key))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n%s\n", jwsPayload, jwsSigned)

	//############

	fmt.Println("JWT & JWS undo:")

	jwtPayloadBytes, err := jws.Verify(jwtSigned, jws.WithKey(jwa.HS512(), jws_key))
	if err != nil {
		panic(err)
	}
	jwtPayload := make(map[string]any)
	err = json.Unmarshal(jwtPayloadBytes, &jwtPayload)
	if err != nil {
		panic(err)
	}

	jwsToken, err := jwt.Parse(jwsSigned, jwt.WithKey(jwa.HS512(), jws_key), jwt.WithValidate(true))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n%+v\n", jwtPayload, jwsToken)

	//#############
	//#############

	fmt.Println("JWE:")

	jwePayload := map[string]any{
		"exp":     time.Now().UTC().Add(3600 * time.Second).Unix(),
		"custom1": "9012",
	}
	jwePayloadBytes, err := json.Marshal(jwePayload)
	if err != nil {
		panic(err)
	}

	jweEncrypted, err := jwe.Encrypt(jwePayloadBytes, jwe.WithKey(jwa.A256GCMKW(), jwe_key))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n%s\n", jwePayload, jweEncrypted)

	//#############

	fmt.Println("JWE undo:")

	jwePayloadBytesUndo, err := jwe.Decrypt(jweEncrypted, jwe.WithKey(jwa.A256GCMKW(), jwe_key))
	if err != nil {
		panic(err)
	}
	jwePayloadUndo := make(map[string]any)
	err = json.Unmarshal(jwePayloadBytesUndo, &jwePayloadUndo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", jwePayloadUndo)
}
