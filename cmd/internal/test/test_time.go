package test

import (
	"fmt"
	"log"
	"time"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
	"gonum.org/v1/plot/plotter"
)

func TestTime() {
	var execTimes plotter.XYs
	execIdx := 1

	start := time.Now()
	bobPrivateKey, err := util.GenerateKeyPair(2048)
	if err != nil {
		log.Fatalf("Erro ao gerar chave privada de Bob: %v", err)
	}

	bobPubPEM := util.ExportPublicKeyAsPEM(&bobPrivateKey.PublicKey)
	bobPrivPEM := util.ExportPrivateKeyAsPEM(bobPrivateKey)

	elapsed := time.Since(start)
	fmt.Printf("Tempo para gerar par de chaves de Bob: %v\n", time.Since(start))
	execTimes = append(execTimes, plotter.XY{X: float64(execIdx), Y: float64(elapsed.Milliseconds())})
	execIdx++

	start = time.Now()
	if err != nil {
		log.Fatalf("Erro ao gerar chave privada de Alice: %v", err)
	}
	fmt.Printf("Tempo para gerar par de chaves de Alice: %v\n", time.Since(start))

	message := "Esta é uma mensagem secreta de Bob para Alice."

	start = time.Now()
	signature, err := util.SignMessage(bobPrivPEM, message)
	if err != nil {
		log.Fatalf("Erro ao assinar a mensagem: %v", err)
	}

	elapsed = time.Since(start)
	fmt.Printf("Tempo para assinar a mensagem: %v\n", time.Since(start))
	execTimes = append(execTimes, plotter.XY{X: float64(execIdx), Y: float64(elapsed.Milliseconds())})
	execIdx++

	start = time.Now()
	err = util.VerifySignature(bobPubPEM, message, signature)
	if err != nil {
		log.Fatalf("Erro na verificação da assinatura: %v", err)
	}
	elapsed = time.Since(start)
	fmt.Printf("Tempo para verificar a assinatura: %v\n", time.Since(start))
	execTimes = append(execTimes, plotter.XY{X: float64(execIdx), Y: float64(elapsed.Milliseconds())})
	execIdx++

	aliceMessage := "Mensagem confidencial de Alice para Bob."
	start = time.Now()
	encryptedMessage, err := util.EncryptMessage(bobPubPEM, aliceMessage)
	if err != nil {
		log.Fatalf("Erro ao cifrar a mensagem: %v", err)
	}
	elapsed = time.Since(start)
	fmt.Printf("Tempo para cifrar a mensagem: %v\n", time.Since(start))
	execTimes = append(execTimes, plotter.XY{X: float64(execIdx), Y: float64(elapsed.Milliseconds())})
	execIdx++

	start = time.Now()
	decryptedMessage, err := util.DecryptMessage(bobPrivPEM, []byte(encryptedMessage))
	if err != nil {
		log.Fatalf("Erro ao decifrar a mensagem: %v", err)
	}
	elapsed = time.Since(start)

	fmt.Printf("Tempo para decifrar a mensagem: %v\n", time.Since(start))
	execTimes = append(execTimes, plotter.XY{X: float64(execIdx), Y: float64(elapsed.Milliseconds())})
	execIdx++

	fmt.Printf("Mensagem decifrada por Bob: %s\n", decryptedMessage)
	if err != nil {
		panic(err)
	}

	util.GenerateGraphicTime(execTimes, "grafico-execucao-algoritmo.png")
}
