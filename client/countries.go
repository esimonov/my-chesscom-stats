package client

import (
	"encoding/json"
	"fmt"

	"github.com/esimonov/my-chesscom-stats/model"
)

// GetCountryByURL fetches full country entity from https://api.chess.com/pub/country/{country_code}.
func (c *Client) GetCountryByURL(countryURL string) (*model.Country, error) {
	resp, err := c.internalClient.Get(countryURL)
	if err != nil {
		return nil, fmt.Errorf("Cannot request country: %s", err)
	}

	country := new(model.Country)
	if err := json.NewDecoder(resp.Body).Decode(country); err != nil {
		return nil, fmt.Errorf("Cannot decode request body: %s", err)
	}

	return country, nil
}
