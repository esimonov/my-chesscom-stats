package client

import (
	"encoding/json"
	"fmt"

	"github.com/esimonov/my-chesscom-stats/model"
)

// GetPlayerByUsername returns chess.com user from https://api.chess.com/pub/player/{username}.
func (c *Client) GetPlayerByUsername(username string) (*model.Player, error) {
	resp, err := c.internalClient.Get(fmt.Sprintf("https://api.chess.com/pub/player/%s", username))
	if err != nil {
		return nil, fmt.Errorf("Cannot request player: %s", err)
	}

	player := new(model.Player)
	if err := json.NewDecoder(resp.Body).Decode(player); err != nil {
		return nil, fmt.Errorf("Cannot decode request body: %s", err)
	}

	return player, nil
}

