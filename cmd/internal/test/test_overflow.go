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

	bobPublicKey := &bobPrivateKey.PublicKey

	util.MeasureAverageTime("geração de chaves de Alice", iterations, func() error {
		_, err := util.GenerateKeyPair(2048)
		return err
	})

	var signature []byte
	util.MeasureAverageTime("assinatura da mensagem por Bob", iterations, func() error {
		var err error
		signature, err = util.SignMessage(bobPrivateKey, message)
		return err
	})

	util.MeasureAverageTime("verificação da assinatura por Alice", iterations, func() error {
		return util.VerifySignature(bobPublicKey, message, signature)
	})

	var encryptedMessage []byte
	util.MeasureAverageTime("cifração da mensagem por Alice", iterations, func() error {
		var err error
		encryptedMessage, err = util.EncryptMessage(bobPublicKey, message)
		return err
	})

	util.MeasureAverageTime("decifração da mensagem por Bob", iterations, func() error {
		_, err := util.DecryptMessage(bobPrivateKey, encryptedMessage)
		return err
	})
}
