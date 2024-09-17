package main

import (
	"fmt"
	"os"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal/chat"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/config"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/service"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/test"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/util"
	"github.com/resend/resend-go/v2"
)

func main() {
	if err := config.LoadEnvs(); err != nil {
		panic(err)
	}

	messageArgs := os.Args[1:]
	message := util.ToString(messageArgs)

	apiKeyResend := config.GetEnv("API_KEY_RESEND")
	client := resend.NewClient(apiKeyResend)

	resendEmail := service.NewResendEmail(client)
	handlerChat := chat.NewChat(resendEmail)

	publicKey, privatekey, err := handlerChat.SendPublicKey()
	if err != nil {
		panic(err)
	}

	signature, cipherMessage, err := handlerChat.SendMessage(message, privatekey, publicKey)
	if err != nil {
		panic(err)
	}

	err = handlerChat.ReceiveMessage(cipherMessage, publicKey, signature, privatekey)
	if err != nil {
		panic(err)
	}

	fmt.Println("------------------------------------------------------------------------------------------------------------------")

	fmt.Println("[TEMPO MEDIO]:")
	test.TestTime()

	fmt.Println("\n[TEMPO MEDIO EXECUTADO 10 VEZES]:")
	test.TestOverflow()
}
