package main

import (
	"github.com/AndrewKraevskii/video-requester/twitch/chat"
)

const client_id = "to0d2ggvuyjpadj2cdxxoxf0kw2flj"
const chat_access = "6bi88uoepkv3bh5ft4vtzy2dz9fvor"

func main() {
	chat.Start(chat_access)
}
