package test

import (
	"crypto/rsa"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
)

const iterations = 10

func TestOverflow() {
	message := "Esta é uma mensagem secreta de Bob para Alice."

	var bobPrivateKey *rsa.PrivateKey

	util.MeasureAverageTime("geração de chaves de Bob", iterations, func() error {
		var err error
		bobPrivateKey, err = util.GenerateKeyPair(2048)
		return err
	})

	bobPubPEM := util.ExportPublicKeyAsPEM(&bobPrivateKey.PublicKey)
	bobPrivPEM := util.ExportPrivateKeyAsPEM(bobPrivateKey)

	util.MeasureAverageTime("geração de chaves de Alice", iterations, func() error {
		_, err := util.GenerateKeyPair(2048)
		return err
	})

	var signature []byte
	util.MeasureAverageTime("assinatura da mensagem por Bob", iterations, func() error {
		var err error
		signature, err = util.SignMessage(bobPrivPEM, message)
		return err
	})

	util.MeasureAverageTime("verificação da assinatura por Alice", iterations, func() error {
		return util.VerifySignature(bobPubPEM, message, signature)
	})

	var encryptedMessage string
	util.MeasureAverageTime("cifração da mensagem por Alice", iterations, func() error {
		var err error
		encryptedMessage, err = util.EncryptMessage(bobPubPEM, message)
		return err
	})

	util.MeasureAverageTime("decifração da mensagem por Bob", iterations, func() error {
		_, err := util.DecryptMessage(bobPrivPEM, []byte(encryptedMessage))
		return err
	})
}
