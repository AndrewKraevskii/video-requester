package pubsub

type WSEvent struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	Timestamp  string     `json:"timestamp"`
	Redemption Redemption `json:"redemption"`
}

type Redemption struct {
	ID         string `json:"id"`
	User       User   `json:"user"`
	ChannelID  string `json:"channel_id"`
	RedeemedAt string `json:"redeemed_at"`
	Reward     Reward `json:"reward"`
	UserInput  string `json:"user_input"`
	Status     string `json:"status"`
}

type Reward struct {
	ID                                string       `json:"id"`
	ChannelID                         string       `json:"channel_id"`
	Title                             string       `json:"title"`
	Prompt                            string       `json:"prompt"`
	Cost                              int64        `json:"cost"`
	IsUserInputRequired               bool         `json:"is_user_input_required"`
	IsSubOnly                         bool         `json:"is_sub_only"`
	Image                             Image        `json:"image"`
	DefaultImage                      Image        `json:"default_image"`
	BackgroundColor                   string       `json:"background_color"`
	IsEnabled                         bool         `json:"is_enabled"`
	IsPaused                          bool         `json:"is_paused"`
	IsInStock                         bool         `json:"is_in_stock"`
	MaxPerStream                      MaxPerStream `json:"max_per_stream"`
	ShouldRedemptionsSkipRequestQueue bool         `json:"should_redemptions_skip_request_queue"`
}

type Image struct {
	URL1X string `json:"url_1x"`
	URL2X string `json:"url_2x"`
	URL4X string `json:"url_4x"`
}

type MaxPerStream struct {
	IsEnabled    bool  `json:"is_enabled"`
	MaxPerStream int64 `json:"max_per_stream"`
}

type User struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
}

