package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"
)

func ExportPublicKeyAsPEM(pubkey *rsa.PublicKey) string {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(pubkey)
	pemPubKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pemPubKey)
}

func GenerateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func SignMessage(privateKey *rsa.PrivateKey, message string) ([]byte, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func VerifySignature(publicKey *rsa.PublicKey, message string, signature []byte) error {
	hashed := sha256.Sum256([]byte(message))
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		return fmt.Errorf("assinatura inválida: %v", err)
	}
	return nil
}

func EncryptMessage(publicKey *rsa.PublicKey, message string) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, []byte(message), label)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func DecryptMessage(privateKey *rsa.PrivateKey, ciphertext []byte) (string, error) {
	label := []byte("")
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func MeasureAverageTime(operationName string, iterations int, operation func() error) time.Duration {
	var totalDuration time.Duration
	for i := 0; i < iterations; i++ {
		start := time.Now()
		err := operation()
		if err != nil {
			log.Fatalf("Erro na operação %s: %v", operationName, err)
		}
		totalDuration += time.Since(start)
	}
	averageDuration := time.Duration(totalDuration.Milliseconds()/int64(iterations)) * time.Millisecond
	fmt.Printf("Tempo médio para %s: %v ms\n", operationName, averageDuration.Milliseconds())

	return averageDuration
}
