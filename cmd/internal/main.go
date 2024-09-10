package internal

import (
	"fmt"
	"log"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
)

func Init() {
	bobPrivateKey, err := util.GenerateKeyPair(2048)
	if err != nil {
		log.Fatalf("Erro ao gerar chave privada de Bob: %v", err)
	}
	bobPublicKey := &bobPrivateKey.PublicKey

	bobPubPEM := util.ExportPublicKeyAsPEM(bobPublicKey)
	fmt.Println("Chave pública de Bob (enviada para Alice):\n", bobPubPEM)

	message := "Esta é uma mensagem secreta de Bob para Alice."

	signature, err := util.SignMessage(bobPrivateKey, message)
	if err != nil {
		log.Fatalf("Erro ao assinar a mensagem: %v", err)
	}
	fmt.Println("Bob assinou a mensagem com sucesso!")
	fmt.Printf("[Assinatura (em bytes)]: %x\n\n", signature)

	aliceHasBobPublicKey := bobPublicKey

	err = util.VerifySignature(aliceHasBobPublicKey, message, signature)
	if err != nil {
		log.Fatalf("Falha na verificação da assinatura: %v", err)
	}

	fmt.Println("Alice verificou a autenticidade da mensagem com sucesso!")
	fmt.Printf("[Mensagem recebida]: %s\n\n", message)
}
