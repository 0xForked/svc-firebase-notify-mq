package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/option"
)

var fcmConn *firebase.App

func (config AppConfig) SetupFCMConnection() {
	opt := option.WithCredentialsFile("firebase.json.example")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Sprintf(
			"failed to initializing firebase app, cause: %s",
			err.Error(),
		))
	}

	setFCMConnection(app)
}

func setFCMConnection(fcmCurrentConn *firebase.App) {
	fcmConn = fcmCurrentConn
}

func (config AppConfig) GetFCMConnection() *firebase.App {
	return fcmConn
}