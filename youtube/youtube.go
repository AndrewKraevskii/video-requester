package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

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
	ContentDetails ContentDetails `json:"contentDetails"`
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

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

func getIdFromText(text string) (string, error) {
	reg := regexp.MustCompile(`((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?`)

	result := reg.FindStringSubmatch(text)
	if len(result) <= 5 {
		return "", errors.New("no youtube url in text")
	}
	value := result[5]
	return value, nil
}

const key = "AIzaSyBVSkJVDJ94CTrr6W4Evpr2ezuTb5TCxYQ"

func getVideoDurationString(videoId string) (string, error) {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/videos?part=contentDetails&id=%s&key=%s", videoId, key)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	parced := Response{}
	err = json.Unmarshal(body, &parced)
	if err != nil {
		return "", err
	}
	return parced.Items[0].ContentDetails.Duration, nil
}

func GetVideoDuration(text string) (time.Duration, error) {
	id, err := getIdFromText(text)
	if err != nil {
		return 0, err
	}
	durationString, err := getVideoDurationString(id)
	if err != nil {
		return 0, err
	}

	duration, err := parceIso(durationString)
	if err != nil {
		return 0, err
	}
	return duration, nil
}
