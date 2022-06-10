package chat

import (
	"github.com/gempir/go-twitch-irc/v3"
)

func Start(access_token string) {
	client := twitch.NewClient("andrewkraevskii", "oauth:"+access_token)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

	})

	client.Join("andrewkraevskii")
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
