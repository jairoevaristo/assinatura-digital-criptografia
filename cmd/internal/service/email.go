package service

import (
	"fmt"

	"github.com/resend/resend-go/v2"
)

type ResendEmail struct {
	client *resend.Client
}

func NewResendEmail(client *resend.Client) *ResendEmail {
	return &ResendEmail{
		client: client,
	}
}

func (r *ResendEmail) Send(to []string, attachments []byte) error {
	params := &resend.SendEmailRequest{
		From:    "Chatbot <onboarding@resend.dev>",
		To:      to,
		Html:    "<p>Olá, essa é sua chave publica para troca de mensagens</p>",
		Subject: "Sua chave pública chegou!",
		Attachments: []*resend.Attachment{
			{
				Content:  attachments,
				Filename: "public-key.txt",
			},
		},
	}

	sent, err := r.client.Emails.Send(params)

	if err != nil {
		return err
	}

	fmt.Println("Email successfully send", sent.Id)
	return nil
}
