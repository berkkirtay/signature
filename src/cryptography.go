package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"time"
)

const (
	RSA_KEY_SIZE = 8192
)

func generateKeyPair() (string, string) {
	keyPair, err := rsa.GenerateKey(rand.Reader, RSA_KEY_SIZE)
	if err != nil {
		panic(err)
	}
	err = keyPair.Validate()
	if err != nil {
		panic(err)
	}
	privateKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(keyPair),
		})

	publicKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&keyPair.PublicKey),
		})

	return string(privateKey), string(publicKey)
}

func generateSignature(privateKey string, data string) string {
	decodedBlock, _ := pem.Decode([]byte(privateKey))
	key, err := x509.ParsePKCS1PrivateKey(decodedBlock.Bytes)
	if err != nil {
		panic(err)
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	hash := generateSHA256Object([]string{data, timestamp})
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		key,
		crypto.SHA256,
		hash[:],
	)
	if err != nil {
		panic(err)
	}
	headers := make(map[string]string)
	headers["timestamp"] = timestamp
	signatureResult := pem.EncodeToMemory(
		&pem.Block{
			Type:    "SIGNATURE",
			Headers: headers,
			Bytes:   signature,
		})
	return string(signatureResult)
}

func verifySignature(data string, signature string, publicKey string) bool {
	decodedSignatureBlock, _ := pem.Decode([]byte(signature))
	decodedPublicKeyBlock, _ := pem.Decode([]byte(publicKey))
	key, err := x509.ParsePKCS1PublicKey(decodedPublicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	hash := generateSHA256Object([]string{data, decodedSignatureBlock.Headers["timestamp"]})
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hash[:], decodedSignatureBlock.Bytes)
	return err == nil
}

func generateSHA256Object(values []string) [32]byte {
	var data []byte = []byte{}
	for _, value := range values {
		data = append(data, []byte(value)...)
	}
	return sha256.Sum256(data)
}

// block := &pem.Block{
// 	Type: "MESSAGE",
// 	Bytes: pubkey_bytes ,
// }

// if err := pem.Encode(os.Stdout, block); err != nil {
// 	log.Fatal(err)
// }
