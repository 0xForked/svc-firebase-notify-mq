package listeners

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/aasumitro/svc-firebase-notify/config"
)

type SingleMessageListener struct {
	token string
	title string
	description string
	config *config.AppConfig
}

// InitMessagingListener initialize the app configuration
func InitSingleMessageListener(
	token string,
	title string,
	description string,
	appConfig *config.AppConfig,
) *SingleMessageListener {
	return &SingleMessageListener{
		token: token,
		title: title,
		description: description,
		config: appConfig,
	}
}

func (listener SingleMessageListener) SendSingleNotify()  {
	fmt.Println(fmt.Sprintf(
		"Send Notify to: %s with title: %s and description: %s",
		listener.token,
		listener.title,
		listener.description,
	))

	ctx := context.Background()
	client, err := listener.config.GetFCMConnection().Messaging(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("error getting Messaging client: %v\n", err))
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := listener.token
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: listener.title,
			Body:  listener.description,
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed send notificaion cause: %v\n", err))
	} else {
		// Response is a message ID string.
		fmt.Println("Successfully sent message:", response)
	}
	fmt.Println("=====================================================")
	// [END send_to_token_golang]
}

