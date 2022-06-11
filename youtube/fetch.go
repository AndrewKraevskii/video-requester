// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func unmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Response) marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Response struct {
	Kind     string   `json:"kind"`
	Etag     string   `json:"etag"`
	Items    []Item   `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

type Item struct {
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
	TopicDetails   TopicDetails   `json:"topicDetails"`
}

type ContentDetails struct {
	Duration        string        `json:"duration"`
	Dimension       string        `json:"dimension"`
	Definition      string        `json:"definition"`
	Caption         string        `json:"caption"`
	LicensedContent bool          `json:"licensedContent"`
	ContentRating   ContentRating `json:"contentRating"`
	Projection      string        `json:"projection"`
}

type ContentRating struct {
}

type Snippet struct {
	PublishedAt          string     `json:"publishedAt"`
	ChannelID            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	Tags                 []string   `json:"tags"`
	CategoryID           string     `json:"categoryId"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
	DefaultLanguage      string     `json:"defaultLanguage"`
	Localized            Localized  `json:"localized"`
	DefaultAudioLanguage string     `json:"defaultAudioLanguage"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Thumbnails struct {
	Default Default `json:"default"`
	Medium  Default `json:"medium"`
	High    Default `json:"high"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type TopicDetails struct {
	TopicCategories []string `json:"topicCategories"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

type VideoNotFound struct{}

func (m *VideoNotFound) Error() string {
	return "video with such id not found"
}

const key = "AIzaSyBVSkJVDJ94CTrr6W4Evpr2ezuTb5TCxYQ"

func FetchVideo(videoId string) (*Item, error) {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/videos?part=contentDetails&part=id&part=liveStreamingDetails&part=snippet&part=topicDetails&id=%s&key=%s", videoId, key)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	response, err := unmarshalResponse(body)
	if err != nil {
		return nil, err
	}
	if response.PageInfo.TotalResults != 1 {
		return nil, &VideoNotFound{}
	}
	return &response.Items[0], nil
}
