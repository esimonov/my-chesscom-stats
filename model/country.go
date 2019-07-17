package model

// Country is how countries are returned by https://api.chess.com/pub/country/{country_code}.
type Country struct {
	APIID string `json:"@id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}
