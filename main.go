package main

import (
	"github.com/gempir/go-twitch-irc/v3"
)

const client_id = "to0d2ggvuyjpadj2cdxxoxf0kw2flj"
const chat_access = "6bi88uoepkv3bh5ft4vtzy2dz9fvor"

func OnPrivateMessage(message twitch.PrivateMessage) {

}

func StartChat(username, access_token string) {
	client := twitch.NewClient(username, access_token)
	client.Join(username)
	client.OnPrivateMessage(OnPrivateMessage)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func main() {

	
	go StartChat(username, access_token)
}
