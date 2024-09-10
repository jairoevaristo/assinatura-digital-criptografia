package test

import (
	"fmt"
	"log"
	"time"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
)

func TestTime() {
	start := time.Now()
	bobPrivateKey, err := util.GenerateKeyPair(2048)
	if err != nil {
		log.Fatalf("Erro ao gerar chave privada de Bob: %v", err)
	}
	bobPublicKey := &bobPrivateKey.PublicKey
	fmt.Printf("Tempo para gerar par de chaves de Bob: %v\n", time.Since(start))

	start = time.Now()
	if err != nil {
		log.Fatalf("Erro ao gerar chave privada de Alice: %v", err)
	}
	fmt.Printf("Tempo para gerar par de chaves de Alice: %v\n", time.Since(start))

	message := "Esta é uma mensagem secreta de Bob para Alice."

	start = time.Now()
	signature, err := util.SignMessage(bobPrivateKey, message)
	if err != nil {
		log.Fatalf("Erro ao assinar a mensagem: %v", err)
	}
	fmt.Printf("Tempo para assinar a mensagem: %v\n", time.Since(start))

	start = time.Now()
	err = util.VerifySignature(bobPublicKey, message, signature)
	if err != nil {
		log.Fatalf("Erro na verificação da assinatura: %v", err)
	}
	fmt.Printf("Tempo para verificar a assinatura: %v\n", time.Since(start))

	aliceMessage := "Mensagem confidencial de Alice para Bob."
	start = time.Now()
	encryptedMessage, err := util.EncryptMessage(bobPublicKey, aliceMessage)
	if err != nil {
		log.Fatalf("Erro ao cifrar a mensagem: %v", err)
	}
	fmt.Printf("Tempo para cifrar a mensagem: %v\n", time.Since(start))

	start = time.Now()
	decryptedMessage, err := util.DecryptMessage(bobPrivateKey, encryptedMessage)
	if err != nil {
		log.Fatalf("Erro ao decifrar a mensagem: %v", err)
	}

	fmt.Printf("Tempo para decifrar a mensagem: %v\n", time.Since(start))
	fmt.Printf("Mensagem decifrada por Bob: %s\n", decryptedMessage)
}
