package main

import (
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/chat"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/config"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/service"
	"github.com/resend/resend-go/v2"
)

func main() {
	if err := config.LoadEnvs(); err != nil {
		panic(err)
	}

	apiKeyResend := config.GetEnv("API_KEY_RESEND")
	client := resend.NewClient(apiKeyResend)

	resendEmail := service.NewResendEmail(client)
	handlerChat := chat.NewChat(resendEmail)

	handlerChat.SendPublicKey([]string{"evaristojairo12@gmail.com"})
	// handlerChat.SendMessage(
	// 	"Isso é uma mensagem secreta",

	// )

	// fmt.Println("[TEMPO MEDIO]:")
	// test.TestTime()

	// fmt.Println("\n[TEMPO MEDIO EXECUTADO 10 VEZES]:")
	// test.TestOverflow()
}
