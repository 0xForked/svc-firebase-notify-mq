package listeners

import "github.com/aasumitro/svc-firebase-notify/config"

type MulticastMessageListener struct {
	tokens []string
	title string
	description string
	config *config.AppConfig
}

// InitMessagingListener initialize the app configuration
func InitMulticastMessageListener(
	tokens []string,
	title string,
	description string,
	appConfig *config.AppConfig,
) *MulticastMessageListener {
	return &MulticastMessageListener{
		tokens: tokens,
		title: title,
		description: description,
		config: appConfig,
	}
}

func (listener MulticastMessageListener) SendMulticastNotify() {
	// TODO
}