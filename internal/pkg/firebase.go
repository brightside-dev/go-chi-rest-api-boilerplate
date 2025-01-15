package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type FirebasePackage struct {
	app *firebase.App
}

func NewFirebaseService(conf *config.Config) *FirebasePackage {
	opt := option.WithCredentialsFile(conf.FirebaseKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	return &FirebasePackage{
		app: app,
	}
}

func (p *FirebasePackage) Send(title, body, token string) {
	// Get the FCM client
	client, err := p.app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Error getting Messaging client: %v", err)
	}

	// Define the message to send
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Hello from FCM!",
			Body:  "This is a test push notification from Go.",
		},
		Token: "your-client-device-token",
	}

	// Send the message
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("Error sending FCM message: %v", err)
	}

	fmt.Printf("Successfully sent message: %s\n", response)
}
