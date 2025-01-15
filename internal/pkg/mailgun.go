package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunClient interface {
	Send(ctx context.Context, message *mailgun.Message) (string, string, error)
}

type MailgunPackage struct {
	client MailgunClient
}

func NewMailgunPackage(client *MailgunClient, conf *config.Config) *MailgunPackage {
	return &MailgunPackage{
		client: mailgun.NewMailgun(conf.MailgunDomain, conf.MailgunAPIKey),
	}
}

func (p *MailgunPackage) Send(sender, subject, body, recipient string) {
	message := mailgun.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := p.client.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
