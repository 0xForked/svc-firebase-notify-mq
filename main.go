package main

import (
	"fmt"
	"github.com/aasumitro/svc-firebase-notify/config"
	"github.com/aasumitro/svc-firebase-notify/src/events"
)

func main() {
	// initialize and setup app configuration
	appConfig := config.InitAppConfig()
	// show current service name and version
	fmt.Println(fmt.Sprintf(
		"Hello world from %s version %s",
		appConfig.GetAppName(),
		appConfig.GetAppVersion(),
	))
	// setup messaging broker/queue connection
	appConfig.SetupAMQPConnection()
	appConfig.SetupFCMConnection()
	// start service events
	fmt.Println("Start listening form messages broker (RabbitMQ)")
	events.InitMessagingEvent(
		appConfig.GetAMQPConnection(),
		appConfig.GetFCMConnection(),
	).ListenToRabbitMQ()
}
