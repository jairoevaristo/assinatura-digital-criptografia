package test

import (
	"crypto/rsa"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
	"gonum.org/v1/plot/plotter"
)

const iterations = 10

var execTimes = make(plotter.XYs, 10)

func TestOverflow() {
	message := "Esta é uma mensagem secreta de Bob para Alice."

	var bobPrivateKey *rsa.PrivateKey

	util.MeasureAverageTime("geração de chaves de Bob", iterations, execTimes, func() error {
		var err error
		bobPrivateKey, err = util.GenerateKeyPair(2048)
		return err
	})

	bobPubPEM := util.ExportPublicKeyAsPEM(&bobPrivateKey.PublicKey)
	bobPrivPEM := util.ExportPrivateKeyAsPEM(bobPrivateKey)

	util.MeasureAverageTime("geração de chaves de Alice", iterations, execTimes, func() error {
		_, err := util.GenerateKeyPair(2048)
		return err
	})

	var signature []byte
	util.MeasureAverageTime("assinatura da mensagem por Bob", iterations, execTimes, func() error {
		var err error
		signature, err = util.SignMessage(bobPrivPEM, message)
		return err
	})

	util.MeasureAverageTime("verificação da assinatura por Alice", iterations, execTimes, func() error {
		return util.VerifySignature(bobPubPEM, message, signature)
	})

	var encryptedMessage string
	util.MeasureAverageTime("cifração da mensagem por Alice", iterations, execTimes, func() error {
		var err error
		encryptedMessage, err = util.EncryptMessage(bobPubPEM, message)
		return err
	})

	util.MeasureAverageTime("decifração da mensagem por Bob", iterations, execTimes, func() error {
		_, err := util.DecryptMessage(bobPrivPEM, []byte(encryptedMessage))
		return err
	})

	util.GenerateGraphicTime(execTimes, "grafico-execucao-algoritmo-10-vezes.png")
}
