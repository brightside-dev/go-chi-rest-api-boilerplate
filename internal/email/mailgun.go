package email

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunService struct {
	Domain string
	APIKey string
}

func NewMailgunService(c *config.Config) *MailgunService {
	return &MailgunService{
		Domain: c.MailgunDomain,
		APIKey: c.MailgunAPIKey,
	}
}

func (s *MailgunService) Send(sender, subject, body, recipient string) {
	mg := mailgun.NewMailgun(s.Domain, s.APIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mailgun.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10-second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
