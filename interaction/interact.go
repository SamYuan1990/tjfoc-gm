package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/Hyperledger-TWGC/tjfoc-gm/x509"
)

func main() {
	privKey, err := x509.ReadPrivateKeyFromPem("priv.pem", nil)
	if err != nil {
		panic(err)
	}
	_, err = x509.ReadCertificateRequestFromPem("pub.pem")
	if err != nil {
		panic(err)
	}
	msg := []byte("abc")
	sign, err := privKey.Sign(rand.Reader, msg, nil) // 签名
	if err != nil {
		log.Fatal(err)
	}
	isok := privKey.PublicKey.Verify(msg, sign)
	fmt.Printf("Verified: %v\n", isok)
}
