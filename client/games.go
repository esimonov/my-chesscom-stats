package client

import (
	"encoding/json"
	"fmt"

	"github.com/esimonov/my-chesscom-stats/model"
)

var (
	monthlyGameArchivesURL = "https://api.chess.com/pub/player/%s/games/archives"
)

// Games are how games are returned by https://api.chess.com/pub/player/{username}/games/{YYYY}/{MM}.
type Games struct {
	Games []model.Game `json:"games"`
}

// GetMonthlyGameArchivesURLList returns list of URLS in the form of https://api.chess.com/pub/player/{username}/games/{YYYY}/{MM}.
func (c *Client) GetMonthlyGameArchivesURLList() ([]string, error) {
	type monthlyGamesArchive struct {
		Archives []string `json:"archives"`
	}

	monthlyGameArchivesURL := fmt.Sprintf(monthlyGameArchivesURL, c.Username)

	resp, err := c.internalClient.Get(monthlyGameArchivesURL)
	if err != nil {
		return nil, fmt.Errorf("Cannot request monthly game archives URL list: %s", err)
	}
	defer resp.Body.Close()

	games := new(monthlyGamesArchive)
	if err := json.NewDecoder(resp.Body).Decode(games); err != nil {
		return nil, fmt.Errorf("Cannot decode request body: %s", err)
	}
	return games.Archives, nil
}

// GetMonthlyGameArchive returns list of URLS in the form of https://api.chess.com/pub/player/{username}/games/{YYYY}/{MM}.
func (c *Client) GetMonthlyGameArchive(url string) ([]model.Game, error) {
	resp, err := c.internalClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot request monthly game archive: %s", err)
	}

	var games Games
	if err := json.NewDecoder(resp.Body).Decode(&games); err != nil {
		return nil, fmt.Errorf("Cannot decode request body: %s", err)
	}

	return games.Games, nil
}
