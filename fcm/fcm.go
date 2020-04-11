package fcm

import (
	"errors"
	"log"

	"github.com/appleboy/go-fcm"
)

// SendPushMessage - send push message
func SendPushMessage(FCMAPIKey string, pushTokens []string, title string, message string, data map[string]interface{}) error {

	if pushTokens == nil || len(pushTokens) == 0 || len(message) == 0 {
		return errors.New("Push token(s) or message is empty")
	}

	notif := &fcm.Notification{
		Body: message,
	}

	//title
	if len(title) > 0 {
		notif.Title = title
	}

	// Create the message to be sent.
	msg := &fcm.Message{
		RegistrationIDs: pushTokens,
		Notification:    notif,
	}

	if data != nil {
		msg.Data = data
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient(FCMAPIKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Send the message and receive the response without retries.
	_, err = client.Send(msg)
	if err != nil {
		log.Println(err)
		return err

	}

	return nil
}
