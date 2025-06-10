package main

import (
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"os"
)

func main() {
	f, err := os.Open("ca.pem")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pemBytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	pemBlock, _ := pem.Decode(pemBytes)
	if pemBlock == nil {
		panic("failed to decode PEM block")
	}

	cert, err := x509.ParseCertificate([]byte(pemBlock.Bytes))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", cert.Raw)

	var asn1Cert x509v3Certificate
	_, err = asn1.Unmarshal(cert.Raw, &asn1Cert)
	if err != nil {
		panic(err)
	}

	certID := pkcs7IssuerAndSerialNumber{
		Issuer:       asn1Cert.TBSCertificate.Issuer,
		SerialNumber: asn1Cert.TBSCertificate.SerialNumber,
	}

	certIDDer, err := asn1.Marshal(certID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", certIDDer)
}

// ASN.1 partial representation of a x.509 v3 certificate
type x509v3Certificate struct {
	TBSCertificate x509TBSCertificate
	// ...
}

// ASN.1 partial representation of the x509TBSCertificate field of a x.509 v3 certificate
type x509TBSCertificate struct {
	Version      int `asn1:"explicit,optional,default:0"`
	SerialNumber *big.Int
	// Signature    asn1.RawValue
	Signature x509AlgorithmIdentifier
	Issuer    x509Name
	// ...
}

// ASN.1 generalization of x509 Name
type x509Name = asn1.RawValue

type x509AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}

// ASN.1 representation of the PKCS #7 pkcs7IssuerAndSerialNumber
type pkcs7IssuerAndSerialNumber struct {
	Issuer       x509Name
	SerialNumber *big.Int
}
