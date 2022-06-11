package main

import (
	"log"
	"sync"

	"github.com/AndrewKraevskii/video-requester/save"
	"github.com/AndrewKraevskii/video-requester/twitch/auth"
	chat "github.com/gempir/go-twitch-irc/v3"
	pubsub "github.com/pajlada/go-twitch-pubsub"
)

const client_id = "gp762nuuoqcoxypju8c569th9wz7q5"

var scopes = []string{
	"chat:read", "chat:edit", "channel:read:redemptions", "channel:manage:redemptions",
}

func GetUser() (*save.User, error) {
	user, err := save.Load()
	if err != nil {
		res, err := auth.GetToken(client_id, scopes)
		if err != nil {
			return nil, err
		}
		info, err := auth.Validate(res.TokenType, res.AccessToken)
		if err != nil {
			return nil, err
		}
		user = &save.User{
			Username:     info.Login,
			Access_token: res.AccessToken,
			Id:           info.UserID,
		}
		err = save.Save(user)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func ChatMessageHandler(message chat.PrivateMessage) {
	log.Println(message.User.DisplayName+":", message.Message)

}

func PointsEventHandler(channelID string, data *pubsub.PointsEvent) {
	log.Println(data.User.DisplayName, "потратил", data.Reward.Cost)
}

func StartChat(user *save.User) {
	chatClient := chat.NewClient(user.Username, "oauth:"+user.Access_token)
	chatClient.Join(user.Username)
	chatClient.OnPrivateMessage(ChatMessageHandler)
	err := chatClient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}

func StartPoints(user *save.User) {
	client := pubsub.NewClient(pubsub.DefaultHost)
	client.Listen(pubsub.PointsEventTopic(user.Id), user.Access_token)
	client.OnPointsEvent(PointsEventHandler)
	err := client.Start()
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	user, err := GetUser()
	if err != nil {
		log.Fatal(err)
	}
	// rw, err := reward.NewClient(user, client_id)
	// _, err = rw.CreateCustomRewards("enabled", 1, "3.1415926")

	client, err := helix.NewClient(&helix.Options{
		ClientID:        client_id,
		UserAccessToken: user.Access_token,
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.UpdateCustomRewards(&helix.UpdateCustomRewardsParams{
		Id:            "07a3eda6-35e5-496b-a333-c2089a08f151",
		BroadcasterID: user.Id,
		Title:         "updated",
		Cost:          1,
		Prompt:        "description changed",
		IsEnabled:     false,
	})

	// res, err := client.CreateCustomReward(&helix.ChannelCustomRewardsParams{
	// 	BroadcasterID: user.Id,
	// 	Title:         "created",
	// 	IsEnabled:     true,
	// 	Cost:          1,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Print(res)
	// res, err := client.GetCustomRewards(&helix.GetCustomRewardsParams{
	// 	BroadcasterID:         user.Id,
	// 	OnlyManageableRewards: true,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, v := range res.Data.ChannelCustomRewards {
	// 	client.DeleteCustomRewards(&helix.DeleteCustomRewardsParams{
	// 		BroadcasterID: user.Id,
	// 		ID: v.ID,
	// 	})
	// }
	// log.Println(res.Data.ChannelCustomRewards)
	// // client.CreateCustomReward(&helix.ChannelCustomRewardsParams{
	// // 	BroadcasterID: user.Id,
	// // 	Title:         "test reward",
	// // 	Cost:          1,
	// // 	Prompt:        "description",
	// // 	IsEnabled:     true,
	// // })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// wg.Add(1)
	// go StartChat(user)
	// wg.Add(1)
	// go StartPoints(user)

	// wg.Wait()
}
