package reward

import (
	"errors"
	"log"

	"github.com/AndrewKraevskii/video-requester/save"
	helix "github.com/nicklaw5/helix"
)

type Client struct {
	hc   *helix.Client
	user *save.User
}

func NewClient(user *save.User, client_id string) (*Client, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        client_id,
		UserAccessToken: user.Access_token,
	})

	return &Client{hc: client, user: user}, err
}

func (client *Client) DeleteCustomRewards(reward_id string) error {
	_, err := client.hc.DeleteCustomRewards(&helix.DeleteCustomRewardsParams{
		BroadcasterID: client.user.Id,
		ID:            reward_id,
	})
	return err
}

func (client *Client) DeleteManageableCustomRewards() error {
	res, err := client.hc.GetCustomRewards(&helix.GetCustomRewardsParams{
		BroadcasterID:         client.user.Id,
		OnlyManageableRewards: true,
	})
	if err != nil {
		return err
	}
	for _, v := range res.Data.ChannelCustomRewards {
		err := client.DeleteCustomRewards(v.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *Client) CreateCustomRewards(title string, cost int, prompt string) (*helix.ChannelCustomReward, error) {
	resp, err := client.hc.CreateCustomReward(&helix.ChannelCustomRewardsParams{
		BroadcasterID: client.user.Id,
		Title:         title,
		Cost:          cost,
		Prompt:        prompt,
		IsEnabled:     true,
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("CreateCustomReward failed")
	}
	log.Println(len(resp.Data.ChannelCustomRewards))
	return &resp.Data.ChannelCustomRewards[0], nil
}

// TODO: Update, (enable/disable), rewardsPull,
