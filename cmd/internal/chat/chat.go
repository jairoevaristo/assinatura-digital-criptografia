package chat

import (
	"fmt"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/service"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
)

type Chat struct {
	resendEmail *service.ResendEmail
}

func NewChat(resendEmail *service.ResendEmail) *Chat {
	return &Chat{
		resendEmail: resendEmail,
	}
}

func (c *Chat) SendPublicKey() (string, string, error) {
	bobPrivateKey, err := util.GenerateKeyPair(2048)
	if err != nil {
		return "", "", err
	}

	bobPubPEM := util.ExportPublicKeyAsPEM(&bobPrivateKey.PublicKey)
	bobPrivPEM := util.ExportPrivateKeyAsPEM(bobPrivateKey)

	// err = c.resendEmail.Send([]string{"jairoevaristodev@gmail.com"}, []byte(bobPubPEM))
	// if err != nil {
	// 	return "", "", err
	// }

	return bobPubPEM, bobPrivPEM, nil
}

func (c *Chat) SendMessage(message string, privateKey string, publicKey string) ([]byte, string, error) {
	cipherMessage, err := util.EncryptMessage(publicKey, message)
	if err != nil {
		return nil, "", err
	}

	signature, err := util.SignMessage(privateKey, cipherMessage)
	if err != nil {
		return nil, cipherMessage, err
	}

	fmt.Println("Alice assinou a mensagem com sucesso!")
	fmt.Printf("[Assinatura]: %x\n\n", signature)

	return signature, cipherMessage, nil
}

func (c *Chat) ReceiveMessage(message string, publicKey string, signature []byte, privateKey string) error {
	err := util.VerifySignature(publicKey, message, signature)
	if err != nil {
		return err
	}

	fmt.Println("Bob verificou a autenticidade da mensagem com sucesso!")

	messageDesciphered, err := util.DecryptMessage(privateKey, []byte(message))
	if err != nil {
		return err
	}

	fmt.Printf("Bob descifrou a mensagem com sucesso!\n")
	fmt.Printf("[Mensagem recebida]: %s\n\n", messageDesciphered)

	return nil
}
