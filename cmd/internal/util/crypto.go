package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"time"

	"gonum.org/v1/plot/plotter"
)

func ExportPublicKeyAsPEM(pubkey *rsa.PublicKey) string {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(pubkey)
	pemPubKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pemPubKey)
}

func ExportPrivateKeyAsPEM(privkey *rsa.PrivateKey) string {
	pubKeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	pemPubKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
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

func SignMessage(privateKey string, message string) ([]byte, error) {
	privateKeyParseRSA, err := parsePrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPSS(rand.Reader, privateKeyParseRSA, crypto.SHA256, hashed[:], nil)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func VerifySignature(publicKey string, message string, signature []byte) error {
	publicKeyParseRSA, err := parsePublicKeyFromPEM(publicKey)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPSS(publicKeyParseRSA, crypto.SHA256, hashed[:], signature, nil)

	if err != nil {
		return nil
	}
	return nil
}

func EncryptMessage(publicKey string, message string) (string, error) {
	publicKeyParseRSA, err := parsePublicKeyFromPEM(publicKey)
	if err != nil {
		return "", err
	}

	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKeyParseRSA, []byte(message), label)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

func DecryptMessage(privateKey string, ciphertext []byte) (string, error) {
	publicKeyParseRSA, err := parsePrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	label := []byte("")
	hash := sha256.New()

	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, publicKeyParseRSA, ciphertext, label)
	if err != nil {

		return "", err
	}
	return string(plaintext), nil
}

func MeasureAverageTime(operationName string, iterations int, execTimes plotter.XYs, operation func() error) time.Duration {
	var totalDuration time.Duration
	for i := 0; i < iterations; i++ {
		start := time.Now()
		err := operation()
		if err != nil {
			log.Fatalf("Erro na operação %s: %v", operationName, err)
		}
		totalDuration += time.Since(start)
		execTimes[i].X = float64(i + 1)
		execTimes[i].Y = float64(totalDuration.Milliseconds())
	}
	averageDuration := time.Duration(totalDuration.Milliseconds()/int64(iterations)) * time.Millisecond
	fmt.Printf("Tempo médio para %s: %v ms\n", operationName, averageDuration.Milliseconds())

	return averageDuration
}

func parsePrivateKeyFromPEM(privateKeyPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, errors.New("falha ao decodificar o bloco PEM")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func parsePublicKeyFromPEM(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, errors.New("falha ao decodificar o bloco PEM")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
