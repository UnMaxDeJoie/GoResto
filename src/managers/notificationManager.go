package Managers

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
)

// SendNotification envoie une notification à un appareil
func SendNotification(serviceAccountKeyPath, deviceToken, title, body string) error {
	// Initialisation de l'application Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'initialisation de l'application Firebase: %v", err)
	}

	// Initialisation du client de messagerie
	client, err := app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'initialisation du client de messagerie: %v", err)
	}

	// Construction de la notification
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: deviceToken,
	}

	// Envoi de la notification
	_, err = client.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'envoi de la notification: %v", err)
	}

	return nil
}
