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

func (c *Chat) SendPublicKey(to []string) (string, error) {
	bobPrivateKey, err := util.GenerateKeyPair(2048)
	if err != nil {
		return "", err
	}

	bobPubPEM := util.ExportPublicKeyAsPEM(&bobPrivateKey.PublicKey)
	// bobPrivPEM := util.ExportPrivateKeyAsPEM(bobPrivateKey)

	err = c.resendEmail.Send(to, []byte(bobPubPEM))
	if err != nil {
		return "", err
	}

	return bobPubPEM, nil
}

func (c *Chat) SendMessage(message string, privateKey string) ([]byte, error) {
	signature, err := util.SignMessage(privateKey, message)
	if err != nil {
		return nil, err
	}

	fmt.Println("Bob assinou a mensagem com sucesso!")
	fmt.Printf("[Assinatura (em bytes)]: %x\n\n", signature)

	return signature, nil
}

func (c *Chat) ReceiveMessage(message string, publicKey string, signature []byte) error {
	err := util.VerifySignature(publicKey, message, signature)
	if err != nil {
		return err
	}

	fmt.Println("Alice verificou a autenticidade da mensagem com sucesso!")
	fmt.Printf("[Mensagem recebida]: %s\n\n", message)

	return nil
}
