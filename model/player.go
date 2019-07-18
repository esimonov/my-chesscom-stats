package model

// Player is how players are returned by https://api.chess.com/pub/player/{username}.
type Player struct {
	Avatar     string `json:"avatar,omitempty"`
	PlayerID   int    `json:"player_id,omitempty"`
	APIID      string `json:"@id,omitempty"`
	URL        string `json:"url,omitempty"`
	Name       string `json:"name,omitempty"`
	Username   string `json:"username,omitempty"`
	Followers  int    `json:"followers,omitempty"`
	CountryURL string `json:"country,omitempty"`
	LastOnline int    `json:"last_online,omitempty"`
	Joined     int    `json:"joined,omitempty"`
	Status     string `json:"status,omitempty"`
	IsStreamer bool   `json:"is_streamer,omitempty"`
}
